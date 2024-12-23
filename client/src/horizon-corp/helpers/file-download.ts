import UseServer from "../request/server"

export async function downloadFile(url: string, fileName: string): Promise<void> {
  const response = await UseServer.get<Blob>(url, {
    responseType: 'blob',
  })

  const blob = new Blob([response.data], { type: response.headers['content-type'] })
  const downloadUrl = URL.createObjectURL(blob)

  // Generate timestamp
  const timestamp = new Date().toISOString().replace(/[-:T.]/g, '').slice(0, 14) // Format: YYYYMMDDHHMMSS
  const fileNameWithTimestamp = `${fileName.replace(/\.[^/.]+$/, '')}_${timestamp}${fileName.match(/\.[^/.]+$/)?.[0] || ''}`

  const a = document.createElement('a')
  a.href = downloadUrl
  a.download = fileNameWithTimestamp
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(downloadUrl)
}