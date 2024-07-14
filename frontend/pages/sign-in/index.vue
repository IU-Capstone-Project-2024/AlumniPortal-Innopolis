<template>
  <div class="bg-[url(~/assets/inno.png)] bg-cover w-screen h-screen flex items-center justify-center">
    <div class="wrapper bg-white p-10 rounded-lg shadow-lg">
      <FormKit 
        type="form" 
        class="space-y-6" 
        submit-label="Register" 
        @submit="submitHandler"
        :actions="false"
      >
        <h2 class="text-3xl font-bold text-center" style="color: #40BA21">Sign in</h2>
        <FormKit 
          type="text" 
          name="login" 
          id="login" 
          validation="required" 
          label="Login" 
          placeholder="Login"
          input-class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
          label-class="block text-sm font-medium text-gray-700" 
        />
        <FormKit 
          type="password" 
          name="password" 
          id="password" 
          validation="required" 
          label="Password"
          placeholder="Password"
          input-class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
          label-class="block text-sm font-medium text-gray-700" 
        />
        <FormKit
            type="checkbox"
            label="Remember me"
            name="remember"
            :value="false"
        />
        <button 
          type="submit" 
          class=" mb-5 w-full flex justify-center py-3 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
          style="background-color: #40BA21"
        >
          Sign in
        </button>
        <UIButton  
          @click="navigateToSignUp"
          text="Register"
          btn_type='button'
         /> 
      </FormKit>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const submitHandler = () => {
  // handle form submission logic here
}

const navigateToSignUp = () => {
  router.push('/sign-up')
}
</script>

<style lang="sass" scoped>
.dark-mode .wrapper
  background-color: black
</style>
