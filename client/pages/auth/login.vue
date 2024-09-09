<template>
  <!-- <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <Card class="w-full max-w-md">
      <CardHeader>
        <CardTitle class="text-2xl font-bold text-center">Login</CardTitle>
        <CardDescription class="text-center">Enter your credentials to access your account</CardDescription>
      </CardHeader>
      <CardContent>
        <form @submit.prevent="handleSubmit">
          <div class="space-y-4">
            <div class="space-y-2">
              <Label for="email">Email</Label>
              <Input v-model="form.email" type="email" id="email" placeholder="johndoe@example.com" required />
              <span v-if="errors.email" class="text-red-500">{{ errors.email }}</span>
            </div>
            <div class="space-y-2">
              <Label for="password">Password</Label>
              <div class="relative">
                <Input v-model="form.password" :type="showPassword ? 'text' : 'password'" id="password"
                  placeholder="••••••••" required />
                <button type="button" @click="togglePasswordVisibility"
                  class="absolute right-3 top-1/2 transform -translate-y-1/2">
                  <template v-if="showPassword">
                    <EyeOffIcon class="h-4 w-4 text-gray-500" />
                  </template>
<template v-else>
                    <EyeIcon class="h-4 w-4 text-gray-500" />
                  </template>
</button>
</div>
<span v-if="errors.password" class="text-red-500">{{ errors.password }}</span>
</div>
</div>
</form>
</CardContent>
<CardFooter class="flex flex-col space-y-4">
  <Button @click="handleSubmit" class="w-full">Log in</Button>
  <div class="text-sm text-center text-gray-500">
    Don't have an account?
    <NuxtLink to="/signup" class="text-primary hover:underline ml-1">Sign up</NuxtLink>
  </div>
  <NuxtLink to="/forgot-password" class="text-sm text-center text-primary hover:underline">
    Forgot your password?
  </NuxtLink>
</CardFooter>
</Card>
</div> -->
  <div>
    Login
  </div>
</template>

<script lang="ts">
// import { defineComponent, ref } from '@nuxtjs/composition-api';
// import { useVuelidate } from '@vuelidate/core';
// import { required, email } from '@vuelidate/validators';
// import { z } from 'zod';
// import { Button } from '@/components/ui/button';
// import { Input } from '@/components/ui/input';
// import { Label } from '@/components/ui/label';
// import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
// import EyeIcon from 'lucide-vue/EyeIcon.vue';
// import EyeOffIcon from 'lucide-vue/EyeOffIcon.vue';

// const schema = z.object({
//   email: z.string().email('Invalid email').min(5),
//   password: z.string().min(8, 'Password must be at least 8 characters long'),
// });

// export default defineComponent({
//   components: {
//     Button,
//     Input,
//     Label,
//     Card,
//     CardContent,
//     CardDescription,
//     CardFooter,
//     CardHeader,
//     CardTitle,
//     EyeIcon,
//     EyeOffIcon,
//   },
//   setup() {
//     const form = ref({
//       email: '',
//       password: '',
//     });

//     const showPassword = ref(false);

//     const togglePasswordVisibility = () => {
//       showPassword.value = !showPassword.value;
//     };

//     const v$ = useVuelidate(schema, form);

//     const handleSubmit = () => {
//       v$.value.$validate();
//       if (v$.value.$error) {
//         console.log('Form has errors');
//         return;
//       }
//       console.log('Form submitted successfully:', form.value);
//       // Perform login logic here
//     };

//     return {
//       form,
//       showPassword,
//       togglePasswordVisibility,
//       handleSubmit,
//       v$: v$.value,
//       errors: v$.value.$errors,
//     };
//   },
// });
</script>

<style scoped>
/* Add your scoped styles here */
</style>
