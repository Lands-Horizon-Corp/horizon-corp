<template>
  <div>
    <div class="grid w-full max-w-sm items-center gap-1.5">
      <Label for="picture">Picture</Label>
      <Input id="picture" type="file" @change="onFileChange" />
    </div>
    <div v-if="uploadProgress">
      <p>{{ uploadProgress.file_name }}: {{ uploadProgress.progress.toFixed(2) }}%</p>
    </div>
    <Progress :model-value="uploadProgress?.progress ?? 0" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useFile } from '@/composables';
import type { UploadProgress } from '@/types';
import { Input } from '@/components/ui/input'
import { Progress } from "@/components/ui/progress"

const selectedFile = ref<File | null>(null);
const uploadProgress = ref<UploadProgress | null>(null);

const fileRepo = useFile();

const emit = defineEmits<{
  (e: 'progress', progress: UploadProgress): void;
  (e: 'uploadSuccess' | 'uploadError', payload: unknown): void;
}>();

function onFileChange(event: Event) {
  const target = event.target as HTMLInputElement;
  const files = target.files;
  if (files && files.length > 0) {
    selectedFile.value = files[0];
    uploadFile()
  }
}

async function uploadFile() {
  if (!selectedFile.value) return;
  try {
    const result = await fileRepo.uploadFileWithProgress(selectedFile.value, (progress) => {
      uploadProgress.value = progress;
      emit('progress', progress);
    });
    emit('uploadSuccess', result);
  } catch (error) {
    emit('uploadError', error);
  }
}

</script>
