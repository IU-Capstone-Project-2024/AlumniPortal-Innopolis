<script setup>
import {useProjectStore} from "~/stores/projectStore.js";
import {useAuthStore} from "~/stores/authStore.js";
import ProposalForm from "~/components/ProposalForm.vue";

const store = useProjectStore()
const auth = useAuthStore()
const projects = store.getProjects().filter(project => project.user.ID === auth.user.ID)

</script>

<template>
<div class="container flex">
  <div>
    <ProposalForm />
  </div>
  <div class="flex flex-col">
    <div v-for="project in projects">
      <span>{{project.title}}</span>
      <span>{{project.desc}}</span>
      <span>{{project.Status}}</span>
    </div>
  </div>
</div>
</template>

<style scoped lang="sass">

</style>