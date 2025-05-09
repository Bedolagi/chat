
const socket = new WebSocket("ws://localhost:8080/ws");

socket.onmessage = function (event) {
    const msg = document.createElement("li");
    msg.textContent = event.data;
    document.getElementById("chatBox").appendChild(msg);
    scrollToBottom();
};

function sendMessage() {
    const input = document.getElementById("msgInput");
    if (input.value.trim() !== "") {
        socket.send(input.value);
        input.value = "";
    }
}

function scrollToBottom() {
    const chatBox = document.getElementById("chatBox");
    chatBox.scrollTop = chatBox.scrollHeight;
}
