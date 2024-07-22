<template>
  <div class="bg-[url(~/assets/inno.png)] bg-cover w-screen h-screen flex items-center justify-center">
    <div class="wrapper">
      <Form @submit="submitHandler" class="form">
        <h2 class="title font-montserrat">Sign up</h2>

        <div class="field">
          <label for="name" class="label font-ibm">Name</label>
          <Field
            name="name"
            as="input"
            type="text"
            placeholder="Name"
            :class="['input', 'font-ibm', { valid: !errors.name && values.name }]"
          />
          <ErrorMessage name="name" class="error-message font-ibm" />
        </div>

        <div class="field">
          <label for="lastname" class="label font-ibm">Lastname</label>
          <Field
            name="lastname"
            as="input"
            type="text"
            placeholder="Lastname"
            :class="['input', 'font-ibm', { valid: !errors.lastname && values.lastname }]"
          />
          <ErrorMessage name="lastname" class="error-message font-ibm" />
        </div>

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
        <div class="field">
          <label for="role" class="label font-ibm">Role</label>
          <Field
              name="role"
              as="select"
              :class="['select', 'font-ibm', { valid: !errors.role && values.role }]"
          >
            <option value="Student">Student</option>
            <option value="Alumni">Alumnus</option>
          </Field>
          <ErrorMessage name="role" class="error-message font-ibm" />
        </div>

        <div>
          <UIButton
            btn_type="submit"
            text="Sign up"
            class="submit-button"
          />
        </div>
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
  name: yup.string().required('Name is required'),
  lastname: yup.string().required('Lastname is required'),
  login: yup.string().email('Must be a valid email').required('Email is required'),
  password: yup.string().required('Password is required'),
  role: yup.string().required('Role is required')
});

const { handleSubmit, values, errors } = useForm({
  validationSchema,
});

const submitHandler = handleSubmit(async (values) => {
  try {
    const res = await fetch('https://api.alumni-portal.ru/auth/signup', {
      method: 'POST',
      credentials: 'include',
      body: JSON.stringify({
        name: values.name,
        lastName: values.lastname,
        email: values.login,
        password: values.password,
        role: values.role,
      }),
    })

    if (res.ok) {
      navigateToSignIn()
    } else {
      const errorData = await res.json();
      throw new Error(errorData.message || 'Signup failed');
    }
  } catch (error) {
  console.error('Signup error', error);
  throw error;
}
});

const navigateToSignIn = () => {
  router.push('/sign-in');
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

  .input, .select
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

.select 
  appearance: none
  background-color: white
  border: 1px solid #D1D5DB
  border-radius: 0.375rem
  padding: 0.5rem 1rem
  line-height: 1.5
  color: #40BA21
  background-image: url('data:image/svg+xml;charset=US-ASCII,<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 4 5"><path fill="%23999" d="M2 0L0 2h4zm0 5L0 3h4z"/></svg>')
  background-repeat: no-repeat
  background-position: right 0.5rem center
  background-size: 1em
  &:focus 
    border-color: #40BA21
    box-shadow: 0 0 0 1px #40BA21
  &::-ms-expand 
    display: none
  

  

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
