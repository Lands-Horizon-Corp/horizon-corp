<template>
  <div>
    <h1 class="text-3xl font-bold underline">Hello world!</h1>
    <div class="p-4 pt-2">
      <Button>Click me YESs</Button>
    </div>
  </div>
</template>
<script setup lang="ts">
import { useRepository } from './composables/repository';


const index = useRepository('users');

onMounted(async () => {
  const users = await index.getAll();
  console.log(users);
})
</script>