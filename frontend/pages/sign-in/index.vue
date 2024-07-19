<template>
  <div class="bg-[url(~/assets/inno.png)] bg-cover w-screen h-screen flex items-center justify-center">
    <div class="wrapper">
      <Form @submit="submitHandler" class="form">
        <h2 class="title font-montserrat">Sign in</h2>

        <div class="field">
          <label for="login" class="label font-ibm">Email</label>
          <Field
            name="login"
            as="input"
            type="email"
            placeholder="Email"
            :class="['input', 'font-ibm', { valid: !errors.login && values.login }]"
          />
          <ErrorMessage name="login" class="error-message font-ibm" />
        </div>

        <div class="field">
          <label for="password" class="label font-ibm">Password</label>
          <Field
            name="password"
            as="input"
            type="password"
            placeholder="Password"
            :class="['input', 'font-ibm', { valid: !errors.password && values.password }]"
          />
          <ErrorMessage name="password" class="error-message font-ibm" />
        </div>

        <div class="field remember-me-field">
          <Field
            type="checkbox"
            name="remember"
            :value="false"
            class="checkbox"
          />
          <label for="remember" class="label font-ibm">Remember me</label>
        </div>

        <div>
          <UIButton
            btn_type="submit"
            text="Sign in"
            class="submit-button"
          />
        </div>
        <UIButton  
          @click="navigateToSignUp"
          text="Register"
          btn_type='button'
        /> 
      </Form>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { useForm, ErrorMessage, Field } from 'vee-validate';
import * as yup from 'yup';

const router = useRouter();

const validationSchema = yup.object({
  login: yup.string().email('Must be a valid email').required('Email is required'),
  password: yup.string().required('Password is required')
});

const { handleSubmit, values, errors } = useForm({
  validationSchema,
});

const submitHandler = handleSubmit((values) => {
  console.log(values);
});

const navigateToSignUp = () => {
  router.push('/sign-up');
};
</script>

<style scoped lang="sass">
.wrapper
  background-color: white
  padding: 1.75rem
  width: 100%
  max-width: 24rem
  border-radius: 0.5rem
  box-shadow: 0 1rem 1.5rem rgba(0, 0, 0, 0.1)

.form
  display: flex
  flex-direction: column
  gap: 1.5rem

.title
  font-size: 1.875rem
  font-weight: bold
  margin-bottom: 0.25rem
  text-align: center
  color: #40BA21

.field
  display: flex
  flex-direction: column

  .label
    margin-top: 0.25rem
    font-size: 0.875rem
    font-weight: 500
    color: #4A5568

  .input
    margin-top: 0.25rem
    width: 100%
    padding: 0.5rem 1rem
    height: 3rem
    border: 1px solid #D1D5DB
    border-radius: 0.375rem
    box-shadow: 0 0.125rem 0.25rem rgba(0, 0, 0, 0.1)
    &:focus
      border-color: #40BA21
      box-shadow: 0 0 0 1px #40BA21
    &.valid
      border-color: #40BA21

  .error-message
    color: #F56565
    font-size: 0.75rem
    margin-top: 0.25rem

.remember-me-field
  display: flex
  align-items: center
  gap: 0.5rem
  flex-direction: row

.checkbox
  width: 1rem
  height: 1rem

.dark-mode .wrapper
  background-color: black
</style>
