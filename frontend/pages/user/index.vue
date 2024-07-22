<template>
  <div class="p-4 mt-[6rem] container mx-auto">
    <h1 class="text-4xl font-bold text-center font-montserrat" style="color: #40BA21">User Info</h1>
    <div class="mt-7 flex justify-between bg-white rounded-lg p-6">
      <div class="flex flex-col items-center w-1/3">
        <div class="w-64 h-64 bg-gray-200 flex items-center justify-center mb-4">
          <img :src="user.image || 'https://via.placeholder.com/150'" alt="Profile" class="rounded-full"/>
        </div>
        <div class="flex space-x-2">
          <UIButton
              btn_type="button"
              text="Upload image"
              class="text-white py-2 px-4 rounded"
              :style="{ backgroundColor: '#40BA21' }"
          />
        </div>
        <div class="flex space-x-2 mt-2">
          <UIButton
              btn_type="submit"
              text="Save"
              @click="saveProfile"
              class="text-white py-2 px-4 rounded"
              :style="{ backgroundColor: '#40BA21' }"
          />
          <UIButton
              btn_type="submit"
              text="Cancel"
              class="text-white py-2 px-4 rounded"
              :style="{ backgroundColor: '#40BA21' }"
          />
        </div>
      </div>
      <div class="w-2/3 pl-10">
        <h3 class="text-xl font-bold font-ibm mb-4" style="color: #40BA21">Profile Overview</h3>
        <div class="mb-2 flex items-center">
          <label class="block text-gray-700 font-ibm w-1/2">First name</label>
          <input type="text" v-model="user.name" class="font-ibm w-2/3 p-2 border rounded" disabled/>
        </div>
        <div class="mb-2 flex items-center">
          <label class="block text-gray-700 font-ibm w-1/2">Last name</label>
          <input type="text" v-model="user.lastName" class="font-ibm w-2/3 p-2 border rounded" disabled/>
        </div>
        <div class="mb-2 flex items-center">
          <label class="block text-gray-700 font-ibm w-1/2">E-mail</label>
          <input type="email" v-model="user.email" class="font-ibm w-2/3 p-2 border rounded" disabled/>
        </div>
        <div class="mb-2 flex items-center">
          <label class="block text-gray-700 font-ibm w-1/2">Role</label>
          <input type="text" v-model="user.role" class="font-ibm w-2/3 p-2 border rounded" disabled/>
        </div>
        <div class="mb-2 flex items-center">
          <label class="block text-gray-700 font-ibm w-1/2">Specialization</label>
          <input type="text" v-model="user.specialization" class="font-ibm w-2/3 l p-2 border rounded" disabled/>
        </div>
        <h3 class="text-xl font-bold font-ibm mb-4 mt-4" style="color: #40BA21">Activity</h3>
        <div class="mb-2 flex items-center">
          <label class="block text-gray-700 font-ibm w-1/2">Recent donations</label>
          <input type="text" class="font-ibm w-2/3 p-2 border rounded" disabled/>
        </div>
        <div class="mb-2 flex items-center">
          <label class="block text-gray-700 font-ibm w-1/2">Volunteering activities</label>
          <input type="text" class="font-ibm w-2/3 p-2 border rounded" disabled/>
        </div>
        <div class="mb-2 flex items-center">
          <label class="block text-gray-700 font-ibm w-1/2">Upcoming events participating</label>
          <input type="text" class="font-ibm w-2/3 p-2 border rounded" disabled/>
        </div>
        <h3 class="text-xl font-bold font-ibm mb-4 mt-4" style="color: #40BA21">Bio</h3>
        <div class="mb-2 flex items-center">
          <label class="block text-gray-700 "></label>
          <textarea v-model="user.description" class="font-ibm w-full p-2 border rounded"></textarea>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta
({
  middleware
      : 'auth'
})
import {useAuthStore} from "~/stores/authStore.js";

const auth = useAuthStore()
import { reactive } from 'vue';

console.log(auth.user)

const user = reactive({
  name: 'test',
  lastName: 'test',
  email: 'test@test.test',
  role: 'Student',
  specialization: '',
  description: '',
  image: '',
});

function saveProfile() {
  // Save profile logic here
  console.log('Profile saved', user);
}

</script>

<style lang="sass" scoped>
</style>