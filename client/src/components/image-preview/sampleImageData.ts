import { MediaResource } from "@/horizon-corp/types";

export const sampleMediaResourceList: MediaResource[] = [
          {
              id: 2,
              fileName: "sample-image.jpg",
              fileSize: 5242880, // 5 MB in bytes
              fileType: "image/jpeg",
              storageKey: "media/resources/sample-image.jpg",
              url: "https://cdn.example.com/media/resources/sample-image.jpg",
              bucketName: "example-media-bucket",
              createdAt: "2024-10-29T08:45:00Z",
              updatedAt: "2024-10-29T10:20:00Z",
              downloadURL: "https://cdn.example.com/media/resources/sample-image.jpg?download=true"
          },
          {
              id: 3,
              fileName: "music-track.mp3",
              fileSize: 7340032, // 7 MB in bytes
              fileType: "audio/mpeg",
              storageKey: "media/resources/music-track.mp3",
              url: "https://cdn.example.com/media/resources/music-track.mp3",
              bucketName: "example-media-bucket",
              createdAt: "2024-10-28T09:30:00Z",
              updatedAt: "2024-10-28T11:00:00Z",
              downloadURL: "https://cdn.example.com/media/resources/music-track.mp3?download=true"
          },
          {
              id: 4,
              fileName: "document.pdf",
              fileSize: 2097152, // 2 MB in bytes
              fileType: "application/pdf",
              storageKey: "media/resources/document.pdf",
              url: "https://cdn.example.com/media/resources/document.pdf",
              bucketName: "example-media-bucket",
              createdAt: "2024-10-27T07:15:00Z",
              updatedAt: "2024-10-27T08:45:00Z",
              downloadURL: "https://cdn.example.com/media/resources/document.pdf?download=true"
          }
      ];
      