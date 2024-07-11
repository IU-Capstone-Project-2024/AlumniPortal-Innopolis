<template>
    <div class="bg-[url(~/assets/inno.png)] bg-cover w-screen h-screen flex items-center justify-center">
        <div class="wrapper bg-white p-10">
            <FormKit type="form" submit-label="Sign in">
                <FormKit type="text" name="login" id="login" validation="required" label="Login" placeholder="Login" />
                <FormKit type="password" name="password" id="password" validation="required" label="Password"
                    placeholder="Password" />
            </FormKit>
        </div>
    </div>
</template>

<script setup>

</script>

<style lang="sass" scoped>
.dark-mode .wrapper
    background-color: black
</style>