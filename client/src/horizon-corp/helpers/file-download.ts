import UseServer from "../request/server"

/**
 * Downloads a file from the specified URL.
 *
 * @param url - The URL to fetch the file from.
 * @param fileName - (Optional) The desired name for the downloaded file. If not provided, the function will attempt to extract it from the response headers.
 */
export async function downloadFile(url: string, fileName?: string): Promise<void> {
  try {
    const response = await UseServer.get<Blob>(url, {
      responseType: 'blob',
    })

    let finalFileName = fileName
    const contentDisposition = response.headers['content-disposition']
    if (!finalFileName && contentDisposition) {
      const fileNameMatch = contentDisposition.match(/filename[^;=\n]*=((['"]).*?\2|[^;\n]*)/)
      if (fileNameMatch != null && fileNameMatch[1]) {
        finalFileName = fileNameMatch[1].replace(/['"]/g, '')
      }
    }

    if (!finalFileName) {
      finalFileName = 'downloaded-file'
    }
    const mimeType = response.headers['content-type'] || 'application/octet-stream'

    const blob = new Blob([response.data], { type: mimeType })
    const generatedURL = window.URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = generatedURL
    a.download = finalFileName
    document.body.appendChild(a)
    a.click()
    // Clean up
    window.URL.revokeObjectURL(generatedURL)
    a.remove()
  } catch (error) {
    console.error('Error downloading the file:', error)
    alert('Failed to download the file. Please try again.')
  }
}
