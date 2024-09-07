<template>
  <div>
    <div class="grid w-full max-w-sm items-center gap-1.5">
      <Label for="picture">Picture</Label>
      <Input id="picture" type="file" multiple @change="onFileChange" />
    </div>
    <div v-for="(progress, index) in uploadProgress" :key="index">
      <p>{{ progress.file_name }}: {{ progress.progress.toFixed(2) }}%</p>
    </div>
    <Progress v-for="(progress, index) in uploadProgress" :key="index" :model-value="progress.progress" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useFile } from '@/composables';
import type { UploadProgress } from '@/types';

const selectedFiles = ref<File[]>([]);
const uploadProgress = ref<UploadProgress[]>([]);

const fileRepo = useFile();

const emit = defineEmits<{
  (e: 'progress', progress: UploadProgress): void;
  (e: 'uploadSuccess' | 'uploadError', payload: unknown): void;
}>();

function onFileChange(event: Event) {
  const target = event.target as HTMLInputElement;
  const files = target.files;
  if (files && files.length > 0) {
    // Clear previous selections if needed, or append to existing ones
    selectedFiles.value = Array.from(files);
    uploadFiles();
  }
}

async function uploadFiles() {
  try {
    const promises = selectedFiles.value.map(async (file) => {
      return await fileRepo.uploadFileWithProgress(file, (progress) => {
        const existingIndex = uploadProgress.value.findIndex(p => p.file_name === file.name);
        if (existingIndex !== -1) {
          uploadProgress.value[existingIndex] = progress;
        } else {
          uploadProgress.value.push(progress);
        }
        emit('progress', progress);
      });
    });

    const results = await Promise.all(promises);
    results.forEach(result => emit('uploadSuccess', result));
  } catch (error) {
    emit('uploadError', error);
  }
}
</script>
