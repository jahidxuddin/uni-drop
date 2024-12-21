const dropArea = document.getElementById("drop-area");
const fileInput = document.getElementById("file-input");
const fileListSection = document.getElementById("file-list-section");
const fileList = document.getElementById("uploaded-files-list");
const chooseRecipientBtn = document.getElementById("choose-recipient-btn");
const sendFilesBtn = document.getElementById("send-files-btn"); // The "Send files" button

let allFiles = [];
let selectedRecipient = null; // Track selected recipient

fileInput.addEventListener("change", (e) => {
	if (fileInput instanceof HTMLInputElement) {
		addFiles(fileInput.files);
	}
});

["dragenter", "dragover"].forEach((eventName) => {
	dropArea.addEventListener(eventName, (e) => {
		e.preventDefault();
		e.stopPropagation();
		dropArea.classList.add("border-blue-500", "bg-blue-100");
	});
});

["dragleave", "drop"].forEach((eventName) => {
	dropArea.addEventListener(eventName, (e) => {
		e.preventDefault();
		e.stopPropagation();
		dropArea.classList.remove("border-blue-500", "bg-blue-100");
	});
});

dropArea.addEventListener("drop", (e) => {
	e.preventDefault();
	if (fileInput instanceof HTMLInputElement) {
		addFiles(e.dataTransfer.files);
	}
});

function addFiles(newFiles) {
	const currentFiles = Array.from(newFiles);

	const dataTransfer = new DataTransfer();

	allFiles.forEach((file) => dataTransfer.items.add(file));

	currentFiles.forEach((file) => {
		if (!allFiles.some((f) => f.name === file.name)) {
			dataTransfer.items.add(file);
			allFiles.push(file);
		}
	});

	fileInput.files = dataTransfer.files;

	updateFileList(allFiles);
	adjustDropAreaWidth();
}

function updateFileList(files) {
	fileList.innerHTML = "";

	files.forEach((file) => {
		const listItem = document.createElement("li");
		listItem.classList.add(
			"file-item",
			"flex",
			"justify-between",
			"bg-gray-100",
			"p-2",
			"my-2",
			"rounded-md",
			"shadow-sm",
			"transition-colors",
			"hover:bg-blue-100"
		);
		listItem.setAttribute("data-filename", file.name);

		const fileInfo = document.createElement("div");
		fileInfo.classList.add("flex", "space-x-2", "items-center");

		const fileName = document.createElement("span");
		fileName.textContent = file.name;
		fileInfo.appendChild(fileName);

		const fileSize = document.createElement("span");
		fileSize.textContent = `(${(file.size / 1024).toFixed(2)} KB)`;
		fileSize.classList.add("text-sm", "text-gray-500");
		fileInfo.appendChild(fileSize);

		listItem.appendChild(fileInfo);

		const removeButton = document.createElement("button");
		removeButton.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="16" height="16" class="text-white"><path d="M12 10.586l4.95-4.95 1.414 1.414L13.414 12l4.95 4.95-1.414 1.414L12 13.414l-4.95 4.95-1.414-1.414L10.586 12l-4.95-4.95 1.414-1.414L12 10.586z"/></svg>`;
		removeButton.classList.add(
			"text-xs",
			"p-1",
			"ml-4",
			"focus:outline-none"
		);

		removeButton.addEventListener("click", () => {
			allFiles = allFiles.filter((f) => f.name !== file.name);

			updateFileList(allFiles);
			fileInput.files = new DataTransfer().files;

			if (allFiles.length === 0) {
				chooseRecipientBtn.classList.add("hidden");
			}
			adjustDropAreaWidth();
		});

		listItem.appendChild(removeButton);
		fileList.appendChild(listItem);
	});

	if (files.length > 0) {
		chooseRecipientBtn.classList.remove("hidden");
		fileListSection.classList.remove("hidden");
	} else {
		chooseRecipientBtn.classList.add("hidden");
		fileListSection.classList.add("hidden");
	}
}

function adjustDropAreaWidth() {
	if (allFiles.length === 0) {
		dropArea.classList.add("w-full");
		dropArea.classList.remove("md:w-1/2");
	} else {
		dropArea.classList.add("md:w-1/2");
		dropArea.classList.remove("w-full");
	}
}

function showModal() {
	document.getElementById("recipient-modal").classList.remove("hidden");
}

function closeModal() {
	document.getElementById("recipient-modal").classList.add("hidden");
}

function selectRecipient(event) {
	let selectedItem = event.target;
	if (selectedItem.tagName !== "LI") {
		selectedItem = selectedItem.closest("li");
	}

	// Deselect previously selected item if any
	if (selectedRecipient && selectedRecipient !== selectedItem) {
		selectedRecipient.classList.remove("bg-blue-100", "text-blue-600");
	}

	// Toggle selection
	selectedItem.classList.toggle(
		"bg-blue-100",
		!selectedItem.classList.contains("bg-blue-100")
	);
	selectedItem.classList.toggle(
		"text-blue-600",
		!selectedItem.classList.contains("text-blue-600")
	);

	// Set the selected recipient if the item is selected, or null if deselected
	if (selectedItem.classList.contains("bg-blue-100")) {
		selectedRecipient = selectedItem;
		sendFilesBtn.classList.remove("hidden"); // Show "Send files" button when selected
	} else {
		selectedRecipient = null;
		sendFilesBtn.classList.add("hidden"); // Hide "Send files" button when no recipient is selected
	}
}

function sendFiles() {
	document.getElementById("file-form").submit();
	closeModal();
}
