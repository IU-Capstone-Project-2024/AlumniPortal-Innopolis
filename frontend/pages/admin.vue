<script setup lang="ts">
import AdminCard from '../components/AdminCard.vue';

const testUser = {
  image: 'https://via.placeholder.com/150',
  name: 'John',
  lastname: 'Doe',
  email: 'john.doe@example.com',
  role: 'Administrator',
};
</script>

<template>
  <div class="p-4 mt-[6rem]">
    <h1 class="text-4xl font-bold text-center font-montserrat" style="color: #40BA21">Admin</h1>
    <div class="mt-7 container flex flex-wrap gap-5 justify-around mx-auto">
      <AdminCard
        :image="testUser.image"
        :name="testUser.name"
        :lastname="testUser.lastname"
        :email="testUser.email"
        :role="testUser.role"
      />
    </div>
  </div>
</template>
