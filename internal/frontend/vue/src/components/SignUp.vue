<template>
  <div class="logo-wrapper">
    <img class="logo" :src="require('../assets/tb_logo_bw.png')"/>
  </div>

  <form id="signup" @submit.prevent="submitForm">
    <div id="username" class="field">
      <label for="username">Username:</label>
      <input type="text" v-model="formData.username" required>
    </div>
    <div id="email" class="field">
      <label for="email">Email:</label>
      <input type="email" v-model="formData.email" required>
    </div>
    <div>
      <button type="submit">Sign Up</button>
    </div>
  </form>

</template>
<script>
export default {
  name: 'SignUpForm',
  data() {
    return {
      formData: {
        username: '',
        email: '',
      },
    };
  },
  methods: {
    async submitForm() {
      try {
        // Send a POST request to the /account/create endpoint
        const response = await fetch('/api/account/create', {
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
          // Handle signup failure, e.g., show error message
          this.$router.push("/signup")
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
#signup {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 400px;
  height: 250px;
}

#signup > div {
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
