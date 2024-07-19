<template>
  <div class="p-4 mt-[6rem]">
    <h1 class="text-4xl font-bold text-center font-montserrat" style="color: #40BA21">All Projects</h1>
    <div class="mt-7 container flex flex-wrap gap-5 justify-around mx-auto">
      <UICard v-for="project in projects" :key="project.id" :image="project.url" :header="project.title"
        :description="project.desc" :project="project" @clicked="(id) => clickHandle(id)"/>
    </div>
  </div>
</template>



<script setup>
import { useProjectStore } from '/stores/projectStore.js'

const router = useRouter();
const clickHandle = (id) => {
  console.log(id)
  router.push('/projects/' + id)
}
const store = useProjectStore()

const projects = store.getProjects
</script>

<style lang="sass" scoped>


</style>