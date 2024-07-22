<template>
<Form @submit="submitHandler" class="form">
  <h2 class="title font-montserrat mb-6">Submit Your Project Proposal</h2>

  <div class="field mb-4">
    <label for="projectName" class="label font-ibm">Project Name</label>
    <Field
      name="projectName"
      as="input"
      type="text"
      placeholder="Project Name"
      :class="['input', 'font-ibm', { valid: !errors.projectName && values.projectName }]"
    />
    <ErrorMessage name="projectName" class="error-message font-ibm" />
  </div>

  <div class="field mb-4">
    <label for="projectDescription" class="label font-ibm">Project Description</label>
    <Field
      name="projectDescription"
      as="textarea"
      rows="18"
      placeholder="Rules:
1. Submissions must be related to IT, computer science, or technology fields.
2. Projects must have a clear objective, scope, and potential impact.
3. A detailed project description must be provided, including goals, methodology, expected outcomes.
4. Submissions should clearly state the problem the project aims to solve and how it proposes to solve it.
5. Projects requiring specialized equipment or software must detail how these will be procured and used.
6. Projects must comply with all university policies and guidelines.
7. Submissions must adhere to ethical standards, including respect for intellectual property, privacy, and data protection laws.
8. Projects involving human subjects must obtain appropriate ethical approvals."
      :class="['input', 'font-ibm', { valid: !errors.projectDescription && values.projectDescription }]"
    />
    <ErrorMessage name="projectDescription" class="error-message font-ibm" />
  </div>

  <div class="field mb-6">
    <label for="goal" class="label font-ibm">Goal (Amount Needed)</label>
    <Field
      name="goal"
      as="input"
      type="number"
      placeholder="Please enter the amount needed"
      :class="['input', 'font-ibm', 'green-text', { valid: !errors.goal && values.goal }]"
    />
    <ErrorMessage name="goal" class="error-message font-ibm" />
  </div>

  <div class="button-wrapper">
    <UIButton
      btn_type="submit"
      text="Submit Proposal"
      class="w-full py-3 text-xl font-semibold"
    />
  </div>
</Form>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useForm, ErrorMessage, Field } from 'vee-validate';
import * as yup from 'yup';

const validationSchema = yup.object({
  projectName: yup.string().required('Project Name is required'),
  projectDescription: yup.string().required('Project Description is required'),
  goal: yup.number().required('Amount is required').positive('Must be a positive number'),
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
  padding: 2rem;
  width: 100%;
  border-radius: 0.5rem;
  box-shadow: 0 1rem 1.5rem rgba(0, 0, 0, 0.1);
}

.form {
  display: flex;
  flex-direction: column;
}

.title {
  font-size: 1.875rem;
  font-weight: bold;
  color: #40BA21;
  text-align: center;
}

.field {
  display: flex;
  flex-direction: column;

  .label {
    margin-top: 0.25rem;
    font-size: 1rem;
    font-weight: 500;
    color: #4A5568;
  }

  .input {
    margin-top: 0.25rem;
    width: 100%;
    padding: 0.75rem 1rem;
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
    font-size: 0.875rem;
    margin-top: 0.25rem;
  }
}

.green-text {
  color: #40BA21;
}

.button-wrapper {
  display: flex;
  justify-content: center;
}
</style>
