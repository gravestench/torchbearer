<template>
  <form id="login" @submit.prevent="submitForm">
    <div>
      <label for="username">Username:</label>
      <input type="text" id="username" v-model="formData.username" required>
    </div>
    <div>
      <label for="password">Password:</label>
      <input type="password" id="password" v-model="formData.password" required>
    </div>
    <div>
      <button type="submit">Login</button>
    </div>
  </form>
</template>

<script>
export default {
  name: 'LoginForm',
  data() {
    return {
      formData: {
        username: '',
        password: '',
      },
    };
  },
  methods: {
    async submitForm() {
      try {
        // Send a POST request to the /account/login endpoint
        const response = await fetch('/api/account/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(this.formData),
        });

        // Handle the response here, e.g., show success message or redirect
        if (response.ok) {
          console.log("redirecting to home")
          // Redirect to a success page or perform other actions
          this.$router.push("/")
        } else {
          // Handle login failure, e.g., show error message
          this.$router.push("/login")
        }
      } catch (error) {
        // Handle network or other errors
        console.error('Error:', error);
      }
    },
  },
};
</script>

<style scoped>
#login {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
</style>
