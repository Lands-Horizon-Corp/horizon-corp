<template>
  <div>
    <h1 class="text-3xl font-bold underline">Upload File</h1>

    <div class="grid w-full max-w-sm items-center gap-1.5 mt-4">
      <Label for="file">Choose a file</Label>
      <Input id="file" type="file" @change="onFileChange" />
    </div>

    <div v-if="uploadProgress !== null" class="mt-4">
      <p>Uploading: {{ uploadProgress.file_name }}</p>
      <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
        <div class="bg-blue-600 h-2.5 rounded-full" :style="{ width: `${uploadProgress.progress}%` }" />
      </div>
      <p>{{ Math.round(uploadProgress.progress) }}%</p>
    </div>

    <div v-if="fileDetails" class="mt-4">
      <h3>File uploaded successfully:</h3>
      <p><strong>File Name:</strong> {{ fileDetails?.file_name }}</p>
      <p><strong>File Size:</strong> {{ (fileDetails?.file_size / 1024).toFixed(2) }} KB</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useFile } from '@/composables/index';
import type { FileDetails, UploadProgress } from '~/types';
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

const fileDetails = ref<FileDetails | null>(null);
const uploadProgress = ref<UploadProgress | null>(null);

const fileRepo = useFile('api/v1'); // Replace 'files' with the correct endpoint

const onFileChange = async (event: Event) => {
  const target = event.target as HTMLInputElement;
  const file = target.files?.[0];

  if (file) {
    try {
      uploadProgress.value = null;
      fileDetails.value = await fileRepo.uploadFileWithProgress(file, (progress: UploadProgress) => {
        uploadProgress.value = progress;
      });
    } catch (error) {
      console.error('Upload failed:', error);
    }
  }
};
</script>

<style scoped>
/* Add any required styles here */
</style>
