const input = document.getElementById("input")
const inputPreview = document.getElementById("inputPreview")
const outputPreview = document.getElementById("outputPreview")
const colorscheme = document.getElementById("colorscheme")

input.addEventListener("change", inputUpdate)
async function inputUpdate() {
  await updateImages(true)
}

colorscheme.addEventListener("change", colorschemeUpdate)
async function colorschemeUpdate() {
  await updateImages(false)
}

async function updateImages(updateInputImage) {
  const file = input.files[0]
  if (!file) {
    return
  }
  if (updateInputImage) {
    inputPreview.src = URL.createObjectURL(file)
  }
  outputPreview.src = "loading.gif"

  // Convert the image and update the output preview
  // Call the backend
  const formData = new FormData()
  formData.append("img", file)

  const url = `http://localhost:3001/convert?colorscheme=${colorscheme.value}`
  const response = await fetch(url, {
    method: "POST",
    body: formData,
  })
  const outputFile = await response.blob()
  outputPreview.src = URL.createObjectURL(outputFile)
}

updateImages(true)
