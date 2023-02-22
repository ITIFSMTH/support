import { createRouter, createWebHistory } from "vue-router";
import VueCookies from 'vue-cookies'
import jwtDecode from "jwt-decode";


const routes = [
  {
    path: '/login',
    name: 'login',
    meta: {layout: 'auth', auth: false},
    component: () => import('@/views/Login.vue')
  },
  {
    path: '/admin',
    name: 'admin',
    meta: {layout: 'panel', auth: true, forAdmin: true},
    component: () => import('@/views/Admin.vue')
  },
  {
    path: '/chats',
    name: 'chats',
    meta: {layout: 'panel', auth: true},
    component: () => import('@/views/Chats.vue')
  },
  {
    path: '/bots',
    name: 'bots',
    meta: {layout: 'panel', auth: true, forAdmin: true},
    component: () => import('@/views/Bots.vue')
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/chats'
  }
];


const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});


router.beforeEach((to, from, next) => {
  const jwtCookie = VueCookies.get('jwt')
  const jwtDecoded = jwtCookie ? jwtDecode(jwtCookie) : {};
  const requireAuth = to.matched.some(record => record.meta.auth);
  const forAdmin = to.matched.some(record => record.meta.forAdmin);

  
  if((requireAuth && !jwtCookie) || (requireAuth && Date.now() >= jwtDecoded.exp * 1000)){
    return next({name:'login'})
  }
  
  if (!requireAuth && Date.now() < jwtDecoded.exp * 1000) return next({name: 'chats'})

  if (forAdmin && jwtDecoded.role !== "admin") return next({name: 'chats'})

  next();
});

export default router;
