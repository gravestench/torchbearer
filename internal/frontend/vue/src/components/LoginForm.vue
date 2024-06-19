<template>
  <div class="logo-wrapper">
    <img class="logo" :src="require('../assets/tb_logo_bw.png')"/>
  </div>

  <form id="login" @submit.prevent="submitForm">
    <div id="username" class="field">
      <label for="username">Username:</label>
      <input type="text" v-model="formData.username" required>
    </div>
    <div id="password" class="field">
      <label for="password">Password:</label>
      <input type="password" v-model="formData.password" required>
    </div>
    <div>
      <button type="submit">Login</button>
    </div>
    <div>Don't have an account? <a href="/signup">Sign Up!</a></div>
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
  width: 400px;
  height: 250px;
}

#login > div {
  margin: 10px;
}

.field > label {
  display: inline-block;
  width: 25%;
  margin-right: 5%;
  text-align: right;
}

.field > input {
  display: inline-block;
  width: 60%;
}

.logo-wrapper {
  display: block;
  position: absolute;
  width: 100vw;
  height: 100vh;
  margin: 0;
  padding: 0;
}

.logo {
  position:relative;
  display:block;
  margin: 0 auto;
  padding: 0;
  height:100%;
  width: auto;
  opacity: 0.07;
}
</style>
