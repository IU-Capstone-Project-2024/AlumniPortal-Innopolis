<script setup lang="ts">
import { ref } from 'vue';
import AdminCard from '../components/AdminCard.vue';
import AdminProjectCard from '../components/AdminProjectCard.vue';

const currentTab = ref('Profiles');

const testUsers = [
  {
    image: 'https://via.placeholder.com/150',
    name: 'John',
    lastname: 'Doe',
    email: 'john.doe@example.com',
    role: 'Administrator',
  },
  {
    image: 'https://via.placeholder.com/150',
    name: 'Jane',
    lastname: 'Doe',
    email: 'jane.doe@example.com',
    role: 'Manager',
  },
  {
    image: 'https://via.placeholder.com/150',
    name: 'Jim',
    lastname: 'Beam',
    email: 'jim.beam@example.com',
    role: 'Developer',
  },
];

const testProjects = [
  {
    projectName: 'Community Garden',
    projectDescription: 'A project to create a community garden in the neighborhood.',
    creatorEmail: 'creator@example.com',
    projectGoal: 5000,
  },
  {
    projectName: 'Tech Conference',
    projectDescription: 'A project to organize a tech conference.',
    creatorEmail: 'tech.creator@example.com',
    projectGoal: 15000,
  },
  {
    projectName: 'Art Festival',
    projectDescription: 'A project to host an art festival.',
    creatorEmail: 'art.creator@example.com',
    projectGoal: 8000,
  },
];
</script>

<template>
  <div class="p-4 mt-[6rem]">
    <h1 class="text-4xl font-bold text-center font-montserrat" style="color: #40BA21">Admin</h1>
    <div class="flex justify-center mt-6 font-ibm" >
      <button 
        @click="currentTab = 'Profiles'" 
        :class="{'active-tab': currentTab === 'Profiles', 'inactive-tab': currentTab !== 'Profiles'}" 
        class="tab-button">
        Profiles
      </button>
      <button 
        @click="currentTab = 'Projects'" 
        :class="{'active-tab': currentTab === 'Projects', 'inactive-tab': currentTab !== 'Projects'}" 
        class="tab-button">
        Projects
      </button>
    </div>
    <div class="mt-7 container flex flex-wrap gap-5 justify-around mx-auto" v-if="currentTab === 'Profiles'">
      <AdminCard
        v-for="(user, index) in testUsers" 
        :key="index"
        :image="user.image"
        :name="user.name"
        :lastname="user.lastname"
        :email="user.email"
        :role="user.role"
      />
    </div>
    <div class="mt-7 container flex flex-wrap gap-5 justify-around mx-auto" v-if="currentTab === 'Projects'">
      <AdminProjectCard
        v-for="(project, index) in testProjects" 
        :key="index"
        :projectName="project.projectName"
        :projectDescription="project.projectDescription"
        :creatorEmail="project.creatorEmail"
        :projectGoal="project.projectGoal"
      />
    </div>
  </div>
</template>

<style lang="sass" scoped>
.tab-button 
  @apply px-4 py-2 mx-2 border rounded transition-colors duration-300
  outline: none
  transition: transform 0.2s
  &:hover 
    transform: scale(1.05)
  &:focus 
    box-shadow: 0 0 0 3px #C0E859
  

.active-tab 
  background-color: #40BA21
  color: white

.inactive-tab 
  background-color: white
  color: #40BA21
  border-color: #40BA21

</style>
