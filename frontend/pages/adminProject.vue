<script setup lang="ts">
import AdminProjectCard from '../components/AdminProjectCard.vue';

const testProject = {
  projectName: 'Community Garden',
  projectDescription: 'A project to create a community garden in the neighborhood.',
  creatorEmail: 'creator@example.com',
  projectGoal: 5000,
};
</script>

<template>
  <div class="p-4 mt-[6rem]">
    <h1 class="text-4xl font-bold text-center font-montserrat" style="color: #40BA21">Admin Projects</h1>
    <div class="mt-7 container flex flex-wrap gap-5 justify-around mx-auto">
      <AdminProjectCard
        :projectName="testProject.projectName"
        :projectDescription="testProject.projectDescription"
        :creatorEmail="testProject.creatorEmail"
        :projectGoal="testProject.projectGoal"
      />
    </div>
  </div>
</template>