<template>
    <div class="pt-20">
        {{ auth.user }}
    </div>
</template>

<script setup>
import {useAuthStore} from "~/stores/authStore.js";

const auth = useAuthStore()


</script>

<style lang="sass" scoped>

</style>