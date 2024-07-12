<template>
    <div class="logo">
      <img v-if="colorMode.value === 'light'" src="~/assets/image.png" alt="logo" class="h-10">
      <img v-if="colorMode.value === 'dark'" src="~/assets/image-w.png" alt="logo" class="h-10">
    </div>
</template>

<script setup>
const colorMode = useColorMode()
console.log(colorMode.value)
</script>

<style lang="sass" scoped>

</style>