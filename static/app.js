// Handle File Upload
document.getElementById("uploadForm").addEventListener("submit", async function(e) {
  e.preventDefault();
  
  const statusElement = document.getElementById("uploadStatus");
  statusElement.innerText = "Uploading...";
  statusElement.style.color = "blue";

  let formData = new FormData(this);

  try {
    let response = await fetch("/upload", {
      method: "POST",
      body: formData
    });

    if (response.ok) {
      statusElement.innerText = await response.text();
      statusElement.style.color = "green";
      this.reset();
    } else {
      statusElement.innerText = "Upload failed: " + await response.text();
      statusElement.style.color = "red";
    }
  } catch (error) {
    statusElement.innerText = "Upload error: " + error.message;
    statusElement.style.color = "red";
  }

  loadFiles();
});

// Load File List
async function loadFiles() {
  try {
    let response = await fetch("/files");
    if (!response.ok) {
      throw new Error('Failed to fetch files');
    }
    
    let files = await response.json();
    let tbody = document.querySelector("#fileTable tbody");
    tbody.innerHTML = "";

    if (files.length === 0) {
      tbody.innerHTML = `<tr><td colspan="5" style="text-align: center;">No files uploaded yet</td></tr>`;
      return;
    }

    files.forEach(file => {
      let row = document.createElement('tr');
      row.innerHTML = `
        <td>${file.id}</td>
        <td>${file.name}</td>
        <td>${(file.size / 1024).toFixed(2)}</td>
        <td>${new Date(file.uploaded_at).toLocaleString()}</td>
        <td><a href="${file.download}" download="${file.name}">⬇️ Download</a></td>
      `;
      tbody.appendChild(row);
    });
  } catch (error) {
    console.error("Error loading files:", error);
    let tbody = document.querySelector("#fileTable tbody");
    tbody.innerHTML = `<tr><td colspan="5" style="text-align: center; color: red;">Error loading files</td></tr>`;
  }
}

// Load on page start
document.addEventListener('DOMContentLoaded', function() {
  loadFiles();
});