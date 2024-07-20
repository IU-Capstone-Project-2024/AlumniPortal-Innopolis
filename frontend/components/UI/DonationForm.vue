<template>
  <div class="wrapper">
    <Form @submit="submitHandler" class="form grid-cols-1">
      <h2 class="title font-montserrat">Make a Donation</h2>
      
      <div class="field col-span-2 sm:col-span-1">
        <label for="firstName" class="label font-ibm">First Name</label>
        <Field
          name="firstName"
          as="input"
          type="text"
          placeholder="First Name"
          :class="['input', 'font-ibm', { valid: !errors.firstName && values.firstName }]"
        />
        <ErrorMessage name="firstName" class="error-message font-ibm" />
      </div>

      <div class="field col-span-2 sm:col-span-1">
        <label for="lastName" class="label font-ibm">Last Name</label>
        <Field
          name="lastName"
          as="input"
          type="text"
          placeholder="Last Name"
          :class="['input', 'font-ibm', { valid: !errors.lastName && values.lastName }]"
        />
        <ErrorMessage name="lastName" class="error-message font-ibm" />
      </div>

      <div class="field col-span-2 sm:col-span-1">
        <label for="email" class="label font-ibm">E-mail</label>
        <Field
          name="email"
          as="input"
          type="email"
          placeholder="E-mail"
          :class="['input', 'font-ibm', { valid: !errors.email && values.email }]"
        />
        <ErrorMessage name="email" class="font-ibm error-message" />
      </div>

      <div class="field col-span-2 sm:col-span-1">
        <label for="phone" class="label font-ibm">Phone number</label>
        <Field
          name="phone"
          as="input"
          type="tel"
          placeholder="Phone number"
          :class="['input', 'font-ibm', { valid: !errors.phone && values.phone }]"
        />
        <ErrorMessage name="phone" class="font-ibm error-message" />
      </div>

      <div class="field col-span-2 sm:col-span-1">
        <label for="amount" class="label font-ibm">Amount</label>
        <Field
          name="amount"
          as="input"
          type="number"
          placeholder="Please enter your donation amount"
          :class="['input', 'font-ibm', 'green-text', { valid: !errors.amount && values.amount }]"
        />
        <ErrorMessage name="amount" class="error-message font-ibm" />
      </div>

      <div class="field col-span-2 sm:col-span-1">
        <label for="paymentMethod" class="label font-ibm">Choose the payment method</label>
        <Field
          name="paymentMethod"
          as="select"
          :class="['select', 'font-ibm', { valid: !errors.paymentMethod && values.paymentMethod }]"
        >
          <option value="Credit/Debit card">Credit/Debit card</option>
          <option value="PayPal">PayPal</option>
          <option value="Bank Transfer">Bank Transfer</option>
        </Field>
        <ErrorMessage name="paymentMethod" class="error-message font-ibm" />
      </div>

      <div class="field col-span-2 sm:col-span-1 full-width">
        <label for="recurringDonation" class="label font-ibm">Recurring donation</label>
        <Field
          name="recurringDonation"
          as="select"
          :class="['select', 'font-ibm', { valid: !errors.recurringDonation && values.recurringDonation }]"
        >
          <option value="Monthly">Monthly</option>
          <option value="Quarterly">Quarterly</option>
          <option value="Yearly">Yearly</option>
        </Field>
        <ErrorMessage name="recurringDonation" class="error-message font-ibm" />
      </div>

      <div class="button-wrapper">
        <UIButton
          btn_type="submit"
          text="Donate"
        />
      </div>
    </Form>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useForm, ErrorMessage, Field } from 'vee-validate';
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

const { handleSubmit, values, errors } = useForm({
  validationSchema,
});

const submitHandler = handleSubmit((values) => {
  console.log(values);
});
</script>

<style scoped lang="scss">
.wrapper {
  background-color: white;
  padding: 1.75rem;
  width: 100%;
  max-width: 48rem;
  border-radius: 0.5rem;
  box-shadow: 0 1rem 1.5rem rgba(0, 0, 0, 0.1);
}

.form {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 1.5rem;
}

.title {
  grid-column: span 2;
  font-size: 1.875rem;
  font-weight: bold;
  margin-bottom: 0.25rem;
  text-align: center;
  color: #40BA21;
}

.field {
  display: flex;
  flex-direction: column;

  .label {
    margin-top: 0.25rem;
    font-size: 0.875rem;
    font-weight: 500;
    color: #4A5568;
  }

  .input, .select {
    margin-top: 0.25rem;
    width: 100%;
    padding: 0.5rem 1rem;
    height: 3rem;
    border: 1px solid #D1D5DB;
    border-radius: 0.375rem;
    box-shadow: 0 0.125rem 0.25rem rgba(0, 0, 0, 0.1);
    &:focus {
      border-color: #40BA21;
      box-shadow: 0 0 0 1px #40BA21;
    }
    &.valid {
      border-color: #40BA21;
    }
  }

  .error-message {
    color: #F56565;
    font-size: 0.75rem;
    margin-top: 0.25rem;
  }
}

.green-text {
  color: #40BA21;
}

.full-width {
  grid-column: span 2;
}

.select {
  appearance: none;
  background-color: white;
  border: 1px solid #D1D5DB;
  border-radius: 0.375rem;
  padding: 0.5rem 1rem;
  line-height: 1.5;
  color: #40BA21;
  background-image: url('data:image/svg+xml;charset=US-ASCII,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 4 5"><path fill="%23999" d="M2 0L0 2h4zm0 5L0 3h4z"/></svg>');
  background-repeat: no-repeat;
  background-position: right 0.5rem center;
  background-size: 1em;
  &:focus {
    border-color: #40BA21;
    box-shadow: 0 0 0 1px #40BA21;
  }
  &::-ms-expand {
    display: none;
  }
}

.button-wrapper {
  grid-column: span 2;
}

.dark-mode .wrapper {
  background-color: black;
}
</style>
