<template>
    <div class="bg-[url(~/assets/inno.png)] bg-cover w-screen h-screen flex items-center justify-center">
        <div class="wrapper bg-white p-7 rounded-lg shadow-lg">
            <FormKit type="form" class="space-y-6" submit-label="Sign Up" @submit="submitHandler" :actions="false"
                #default="{ value }">
                <h2 class="text-3xl font-bold text-center" style="color: #40BA21">Sign up</h2>
                <FormKit type="text" name="login" id="login" validation="required" label="Login" placeholder="Login"
                    input-class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
                    label-class="block text-sm font-medium text-gray-700" />
                <FormKit type="password" name="password" id="password" validation="required" label="Password"
                    placeholder="Password"
                    input-class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
                    label-class="block text-sm font-medium text-gray-700" />
                <button type="submit"
                    class="w-full flex justify-center py-3 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white hover:bg-green-700; focus:outline-none; focus:ring-2; focus:ring-offset-2; focus:ring-green-500"
                    style="background-color: #40BA21">
                    Sign Up
                </button>
            </FormKit>
        </div>
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router'

const router = useRouter()

const handleSubmit = () => {
    // handle form submission logic here
}

const navigateToSignIn = () => {
    router.push('/sign-in')
}
</script>

<style lang="sass" scoped>
.dark-mode .wrapper
  background-color: black
</style>
