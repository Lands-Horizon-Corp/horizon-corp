<template>
  helo
</template>
<!-- <template>
  <div v-for="(progress, index) in uploadProgresses" :key="index">
    <p>{{ progress.fileName }}: {{ progress.percentage.toFixed(2) }}%</p>
    <button @click="cancelUpload(progress.fileId)">Cancel</button>
    <progress :value="progress.percentage" max="100"></progress>
  </div>
</template>

<script setup lang="ts">
import { ref, onUnmounted } from 'vue';
import axios from 'axios';
import { FileRepository } from '@/composables'; // Adjust the path as necessary
import type { UploadProgress } from '@/types'; // Adjust the path as necessary

const uploadProgresses = ref<UploadProgress[]>([]);
const fileRepository = new FileRepository('https://your-api-url.com');

const uploadFile = async (file: File) => {
  const cancelTokenSource = axios.CancelToken.source();
  const progressHandler = (progress: UploadProgress) => {
    const index = uploadProgresses.value.findIndex(p => p?.fileId === progress.fileId);
    if (index !== -1) {
      uploadProgresses.value[index] = progress;
    } else {
      uploadProgresses.value.push({ ...progress, fileId: file.name });
    }
  };

  try {
    const result = await fileRepository.uploadFileWithProgress(
      file,
      progressHandler,
      cancelTokenSource,
      () => console.log('Upload cancelled')
    );
    console.log('Upload success', result);
  } catch (error) {
    console.error('Upload error', error);
  }
};

const cancelUpload = (fileName: string) => {
  // You need to store and access the cancel token for each file
  // cancelTokenSource.cancel('User cancelled the upload');
};

onUnmounted(() => {
  // Make sure to cancel all ongoing uploads
  // cancelTokens.forEach(token => token.cancel());
});
</script>

<style scoped>
/* Add any styles here */
</style> -->


<!-- <template>
  <div>
    <div class="grid w-full max-w-sm items-center gap-1.5">
      <Label for="picture">Picture</Label>
      <Input id="picture" type="file" multiple @change="onFileChange" />
    </div>
    <div v-for="(progress, index) in uploadProgress" :key="index">
      <p>{{ progress.file_name }}: {{ progress.progress.toFixed(2) }}%</p>
      <button @click="cancelUpload(progress.file_name)">Cancel</button>
      <Progress :model-value="progress.progress" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import axios, { type CancelTokenSource } from 'axios';
import { useFile } from '@/composables';
import type { UploadProgress } from '@/types';

const selectedFiles = ref<File[]>([]);
const uploadProgress = ref<UploadProgress[]>([]);
const cancelTokenSources = ref<Map<string, CancelTokenSource>>(new Map());

const fileRepo = useFile();

const emit = defineEmits<{
  (e: 'progress', progress: UploadProgress): void;
  (e: 'uploadSuccess' | 'uploadError' | 'uploadCancelled', payload: unknown): void;
}>();

function onFileChange(event: Event) {
  const target = event.target as HTMLInputElement;
  const files = target.files;
  if (files && files.length > 0) {
    selectedFiles.value = Array.from(files);
    uploadFiles();
  }
}

function onCancelUpload(fileName: string) {
  emit('uploadCancelled', { fileName });
}

async function uploadFiles() {
  try {
    const promises = selectedFiles.value.map(async (file) => {
      const cancelToken = axios.CancelToken.source();
      cancelTokenSources.value.set(file.name, cancelToken);

      return await fileRepo.uploadFileWithProgress(file, (progress) => {
        const existingIndex = uploadProgress.value.findIndex(p => p.file_name === file.name);
        if (existingIndex !== -1) {
          uploadProgress.value[existingIndex] = progress;
        } else {
          uploadProgress.value.push(progress);
        }
        emit('progress', progress);
      }, cancelToken, () => onCancelUpload(file.name)); // Passing the onCancel callback
    });
    const results = await Promise.all(promises);
    results.forEach(result => emit('uploadSuccess', result));
  } catch (error) {
    emit('uploadError', error);
  }
}

function cancelUpload(fileName: string) {
  const cancelToken = cancelTokenSources.value.get(fileName);
  if (cancelToken) {
    cancelToken.cancel(`Cancelled upload: ${fileName}`);
  }
}
</script> -->
