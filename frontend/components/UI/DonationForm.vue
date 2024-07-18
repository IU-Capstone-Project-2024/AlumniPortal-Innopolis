<template>
  <div class="wrapper bg-white p-7 w-[25rem] rounded-lg shadow-lg">
    <Form @submit="submitHandler" class="space-y-6">
      <h2 class="text-3xl font-bold mb-1 font-montserrat text-center" style="color: #40BA21">Make a Donation</h2>
      
      <div>
        <label for="firstName" class="mt-1 block text-sm font-medium font-ibm text-gray-700">First Name</label>
        <Field
          name="firstName"
          as="input"
          type="text"
          placeholder="First Name"
          class="mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
        <ErrorMessage name="firstName" class="text-red-500 text-xs mt-1" />
      </div>

      <div>
        <label for="lastName" class="block text-sm font-medium font-ibm text-gray-700">Last Name</label>
        <Field
          name="lastName"
          as="input"
          type="text"
          placeholder="Last Name"
          class="mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
        <ErrorMessage name="lastName" class="text-red-500 text-xs mt-1" />
      </div>

      <div>
        <label for="email" class="block text-sm font-medium font-ibm text-gray-700">E-mail</label>
        <Field
          name="email"
          as="input"
          type="email"
          placeholder="E-mail"
          class="mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
        <ErrorMessage name="email" class="text-red-500 text-xs mt-1" />
      </div>

      <div>
        <label for="phone" class="block text-sm font-medium font-ibm text-gray-700">Phone number</label>
        <Field
          name="phone"
          as="input"
          type="tel"
          placeholder="Phone number"
          class="mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
        <ErrorMessage name="phone" class="text-red-500 text-xs mt-1" />
      </div>

      <div>
        <label for="amount" class="block text-sm font-medium font-ibm text-gray-700">Amount</label>
        <Field
          name="amount"
          as="input"
          type="number"
          placeholder="Please enter your donation amount"
          class="mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        />
        <ErrorMessage name="amount" class="text-red-500 text-xs mt-1" />
      </div>

      <div>
        <span class="block text-sm mb-2 italic font-medium font-montserrat text-gray-700">Choose the payment method</span>
        <Field
          name="paymentMethod"
          as="select"
          class="mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        >
          <option value="Credit/Debit card">Credit/Debit card</option>
          <option value="PayPal">PayPal</option>
          <option value="Bank Transfer">Bank Transfer</option>
        </Field>
      </div>

      <div>
        <span class="block text-sm mb-2 italic font-montserrat font-medium text-gray-700">Recurring donation</span>
        <Field
          name="recurringDonation"
          as="select"
          class="mt-1 block w-full p-1 font-ibm border border-gray-300 rounded-md shadow-sm focus:ring-green-500 focus:border-green-500"
        >
          <option value="Monthly">Monthly</option>
          <option value="Quarterly">Quarterly</option>
          <option value="Yearly">Yearly</option>
        </Field>
      </div>

      <UIButton
        btn_type="submit"
        text="Donate"
      />
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

<style lang="sass" scoped>
.dark-mode .wrapper
  background-color: black
</style>
