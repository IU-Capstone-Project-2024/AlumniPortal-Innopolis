<template>
  <div class="wrapper bg-white p-7 w-full max-w-3xl rounded-lg shadow-lg">
    <Form @submit="submitHandler" class="space-y-6 grid grid-cols-2 gap-6">
      <h2 class="col-span-2 text-3xl font-bold mb-1 font-montserrat text-center" style="color: #40BA21">Make a Donation</h2>
      
      <div class="col-span-1">
        <label for="firstName" class="mt-1 block text-sm font-medium font-ibm text-gray-700">First Name</label>
        <Field
          name="firstName"
          as="input"
          type="text"
          placeholder="First Name"
          class="mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
        <ErrorMessage name="firstName" class="text-red-500 font-ibm text-xs mt-1" />
      </div>

      <div class="col-span-1">
        <label for="lastName" class="block text-sm font-medium font-ibm text-gray-700">Last Name</label>
        <Field
          name="lastName"
          as="input"
          type="text"
          placeholder="Last Name"
          class="mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
        <ErrorMessage name="lastName" class="text-red-500 font-ibm text-xs mt-1" />
      </div>

      <div class="col-span-1">
        <label for="email" class="block text-sm font-medium font-ibm text-gray-700">E-mail</label>
        <Field
          name="email"
          as="input"
          type="email"
          placeholder="E-mail"
          class="mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
        <ErrorMessage name="email" class="text-red-500 font-ibm text-xs mt-1" />
      </div>

      <div class="col-span-1">
        <label for="phone" class="block text-sm font-medium font-ibm text-gray-700">Phone number</label>
        <Field
          name="phone"
          as="input"
          type="tel"
          placeholder="Phone number"
          class="mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
        <ErrorMessage name="phone" class="text-red-500 font-ibm text-xs mt-1" />
      </div>

      <div class="col-span-1">
        <label for="amount" class="block text-sm font-medium font-ibm text-gray-700">Amount</label>
        <Field
          name="amount"
          as="input"
          type="number"
          placeholder="Please enter your donation amount"
          class="custom-input mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
        <ErrorMessage name="amount" class="text-red-500 font-ibm text-xs mt-1" />
      </div>

      <div class="col-span-1">
        <label for="paymentMethod" class="block text-sm font-medium font-ibm text-gray-700">Choose the payment method</label>
        <Field
          name="paymentMethod"
          as="select"
          class="custom-select mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        >
          <option value="Credit/Debit card">Credit/Debit card</option>
          <option value="PayPal">PayPal</option>
          <option value="Bank Transfer">Bank Transfer</option>
        </Field>
        <ErrorMessage name="paymentMethod" class="text-red-500 font-ibm text-xs mt-1" />
      </div>

      <div class="col-span-2">
        <label for="recurringDonation" class="block text-sm font-ibm font-medium text-gray-700">Recurring donation</label>
        <Field
          name="recurringDonation"
          as="select"
          class="custom-select mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        >
          <option value="Monthly">Monthly</option>
          <option value="Quarterly">Quarterly</option>
          <option value="Yearly">Yearly</option>
        </Field>
        <ErrorMessage name="recurringDonation" class="text-red-500 font-ibm text-xs mt-1" />
      </div>

      <div class="col-span-2">
        <UIButton
          btn_type="submit"
          text="Donate"
        />
      </div>
    </Form>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useField, useForm, ErrorMessage, Field } from 'vee-validate';
import * as yup from 'yup';

const validationSchema = yup.object({
  firstName: yup.string().required('First Name is required'),
  lastName: yup.string().required('Last Name is required'),
  email: yup.string().email('Must be a valid email').required('E-mail is required'),
  phone: yup.string().required('Phone number is required'),
  amount: yup.number().required('Amount is required').positive('Must be a positive number'),
  paymentMethod: yup.string().required('Payment method is required'),
  recurringDonation: yup.string().required('Recurring donation is required')
});

const { handleSubmit } = useForm({
  validationSchema,
});

const submitHandler = handleSubmit((values) => {
  console.log(values);
});
</script>

<style scoped>
.dark-mode .wrapper {
  background-color: black;
}

.custom-select {
  appearance: none;
  background-color: white;
  border: 1px solid #d1d5db;
  border-radius: 0.375rem;
  padding: 0.5rem 1rem;
  line-height: 1.5;
  color: #40BA21;
  background-image: url('data:image/svg+xml;charset=US-ASCII,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 4 5"><path fill="%23999" d="M2 0L0 2h4zm0 5L0 3h4z"/></svg>');
  background-repeat: no-repeat;
  background-position: right 0.5rem center;
  background-size: 1em;
}

.custom-select::-ms-expand {
  display: none;
}

.custom-select:focus {
  border-color: #40ba21;
  box-shadow: 0 0 0 1px #40ba21;
}

.custom-input {
  color: #40BA21;
  border: 1px solid #d1d5db;
  border-radius: 0.375rem;
  padding: 0.5rem 1rem;
  line-height: 1.5;
}

.custom-input:focus {
  border-color: #40ba21;
  box-shadow: 0 0 0 1px #40ba21;
}

.grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 1.5rem;
}
</style>
