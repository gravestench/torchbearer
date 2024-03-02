import { createRouter, createWebHashHistory } from 'vue-router'
import HelloWorld from "../components/HelloWorld.vue";
import LoginForm from "../components/LoginForm.vue";

const routes = [
  {
    path: '/*',
    beforeEnter: (to, from, next) => {
      if (to.path.startsWith('/api/')) {
        next(false);
      } else {
        next()
      }
    },
  },
  {
    path: '/',
    component: HelloWorld,
    meta: { requiresAuth: true },
  },
  {
    path: '/login',
    component: LoginForm,
  },
]

// Function to check authentication
async function checkAuthentication() {
  try {
    // Send a GET request to the /account/authenticated endpoint using the fetch API
    const response = await fetch(`http://localhost:8080/api/account/authenticated`);

    // Check the HTTP status code
    if (response.ok) {
      // User is authenticated
      return true;
    } else {
      // User is not authenticated
      return false;
    }
  } catch (error) {
    // An error occurred (e.g., network error)
    console.error('Authentication check failed:', error);
    return false;
  }
}

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

// Add a global navigation guard to check for authentication
router.beforeEach((to, from, next) => {
  const isAuthenticated = checkAuthentication();

  // if not requires auth, continue
  if (!to.matched.some(record => record.meta.requiresAuth)) {
    next();
    return
  }

  isAuthenticated.then(val =>{
    // if authenticated, continue
    if (val) {
      next();
      return
    }

    // redirect to login page
    next('/login');
  })
});

export default router
