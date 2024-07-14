<template>
    <div class="wrapper bg-white p-7 w-[25rem] rounded-lg shadow-lg">
      <FormKit 
        type="form" 
        class="space-y-6" 
        submit-label="Donate" 
        @submit="submitHandler"
        :actions="false"
      >
        <h2 class="text-3xl font-bold font-montserrat text-center" style="color: #40BA21">Make a Donation</h2>
        <FormKit 
          type="text" 
          name="firstName" 
          id="firstName" 
          validation="required" 
          label="First Name" 
          placeholder="First Name"
          input-class="mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
          label-class="mt-1 block text-sm font-medium font-ibm text-gray-700" 
        />
        <FormKit 
          type="text" 
          name="lastName" 
          id="lastName" 
          validation="required" 
          label="Last Name"
          placeholder="Last Name"
          input-class="mt-1 font-ibm block w-full p-1 border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
          label-class="block font-ibm text-sm font-medium text-gray-700" 
        />
        <FormKit 
          type="email" 
          name="email" 
          id="email" 
          validation="required|email" 
          label="E-mail"
          placeholder="E-mail"
          input-class="mt-1 font-ibm block w-full p-1 border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
          label-class="block font-ibm text-sm font-medium text-gray-700" 
        />
        <FormKit 
          type="tel" 
          name="phone" 
          id="phone" 
          validation="required|tel" 
          label="Phone number"
          placeholder="Phone number"
          input-class="mt-1 font-ibm block w-full p-1 border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
          label-class="block font-ibm text-sm font-medium text-gray-700" 
        />
        <FormKit 
          type="number" 
          name="amount" 
          id="amount" 
          validation="required|number" 
          label="Amount"
          placeholder="Please enter your donation amount"
          input-class="mt-1 font-ibm block w-full p-1 border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
          label-class="block font-ibm text-sm font-medium text-gray-700" 
        />
        <div>
          <span class="block text-sm mb-1 font-medium font-montserrat text-gray-700">Choose the payment method</span>
          <FormKit 
            type="radio" 
            name="paymentMethod" 
            v-model="form.paymentMethod" 
            :options="paymentMethods"
            input-class="mr-2"
            label-class="text-sm font-ibm font-medium text-gray-700"
          />
        </div>
        <div>
          <span class="block text-sm mb-1 font-ibm font-medium text-gray-700">Recurring donation</span>
          <FormKit 
            type="radio" 
            name="recurringDonation" 
            v-model="form.recurringDonation" 
            :options="recurringDonations"
            input-class="mr-2"
            label-class="text-sm font-ibm font-medium text-gray-700"
          />
        </div>
        <UIButton
          btn_type="submit"
          text="Donate"
        />
      </FormKit>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const form = ref({
  paymentMethod: 'Credit/Debit card',
  recurringDonation: 'Monthly',
});

const paymentMethods = [
  { value: 'Credit/Debit card', label: 'Credit/Debit card' },
  { value: 'PayPal', label: 'PayPal' },
  { value: 'Bank Transfer', label: 'Bank Transfer' },
];

const recurringDonations = [
  { value: 'Monthly', label: 'Monthly' },
  { value: 'Quarterly', label: 'Quarterly' },
  { value: 'Yearly', label: 'Yearly' },
];

const submitHandler = (formData) => {
  console.log(formData);
};
</script>

<style lang="sass" scoped>
.dark-mode .wrapper
  background-color: black
</style>
