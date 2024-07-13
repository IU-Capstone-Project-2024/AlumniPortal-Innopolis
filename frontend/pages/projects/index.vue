<template>
  <div class="p-4 mt-[6rem]">
    <h1 class="text-3xl font-bold text-center font-montserrat" style="color: #40BA21">All Projects</h1>
    <div class="mt-7 container flex flex-wrap gap-5 justify-around mx-auto">
      <UICard v-for="project in projects" :key="project.id" :image="project.url" :header="project.title"
        :description="project.desc" />
    </div>
  </div>
</template>



<script setup>
  const projects = [
    {
      id: 1,
      url: '_nuxt/assets/inno.png',
      title: 'Inno',
      desc: 'description'
    },
    {
      id: 2,
      url: '_nuxt/assets/inno.png',
      title: 'Inno',
      desc: 'description'
    },
    {
      id: 3,
      url: '_nuxt/assets/inno.png',
      title: 'Inno',
      desc: 'description'
    },
    {
      id: 4,
      url: '_nuxt/assets/inno.png',
      title: 'Inno',
      desc: 'description'
    },
    {
      id: 5,
      url: '_nuxt/assets/inno.png',
      title: 'Inno',
      desc: 'description'
    },
    {
      id: 6,
      url: '_nuxt/assets/inno.png',
      title: 'Inno',
      desc: 'description'
    },
    {
      id: 7,
      url: '_nuxt/assets/inno.png',
      title: 'Inno',
      desc: 'description'
    },
    {
      id: 8,
      url: '_nuxt/assets/inno.png',
      title: 'Inno',
      desc: 'description'
    },
    {
      id: 9,
      url: '_nuxt/assets/inno.png',
      title: 'Inno',
      desc: 'description'
    },
    {
      id: 10,
      url: '_nuxt/assets/inno.png',
      title: 'Inno',
      desc: 'description'
    },
  ]
</script>

<style lang="sass" scoped>

</style>