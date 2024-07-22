import {useAuthStore} from "~/stores/authStore";

export default defineNuxtRouteMiddleware((to, from) => {
    const auth = useAuthStore()
    console.log(auth.user)
    if (!auth.user) {
        return navigateTo('/sign-in')
    }
})
