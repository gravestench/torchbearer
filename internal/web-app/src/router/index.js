import { createRouter, createWebHistory } from 'vue-router'
import HomePage from "../components/HomePage.vue";
import LoginForm from "../components/LoginForm.vue";

const routes = [
  {
    path: '/login',
    component: LoginForm,
  },
  {
    path: '/',
    component: HomePage,
    meta: { requiresAuth: true },
  },
  {
    path: '/api/*',
    beforeEnter: (to, from, next) => {
      next(false)
    },
  },
]

// Function to check authentication
async function checkAuthentication() {
  try {
    // Send a GET request to the /account/authenticated endpoint using the fetch API
    const response = await fetch(`/api/account/authenticated`);
    console.log(response)

    // Check the HTTP status code
    if (response.ok) {
      // User is authenticated
      console.log("authenticated")
      return true;
    } else {
      // User is not authenticated
      console.log("not authenticated")
      return false;
    }
  } catch (error) {
    // An error occurred (e.g., network error)
    console.log("error authenticated")
    console.error('Authentication check failed:', error);
    return false;
  }
}

const router = createRouter({
  history: createWebHistory(),
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
      console.log(to.path)
      // redirect authenticated users away from login page
      if (to.path === "/login") {
        next("/");
        return
      }

      next();
      return
    }

    // redirect to login page
    next('/login');
  })
});

export default router
