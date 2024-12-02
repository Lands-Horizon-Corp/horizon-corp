import UseServer from "../request/server"

export async function downloadFile(url: string, fileName: string): Promise<void> {
  const response = await UseServer.get<Blob>(url, {
    responseType: 'blob',
  })

  const blob = new Blob([response.data], { type: response.headers['content-type'] })
  const downloadUrl = URL.createObjectURL(blob)

  const a = document.createElement('a')
  a.href = downloadUrl
  a.download = fileName
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(downloadUrl)
}