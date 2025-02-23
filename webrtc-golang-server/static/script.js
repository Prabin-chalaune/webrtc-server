const localVideo = document.getElementById('localVideo');
const remoteVideo = document.getElementById('remoteVideo');
const ws = new WebSocket('ws://localhost:8000/ws');

// Create a new RTCPeerConnection
const pc = new RTCPeerConnection({
    iceServers: [{ urls: 'stun:stun.l.google.com:19302' }]
});

// Handling incoming tracks
pc.ontrack = (event) => {
    if (event.track.kind === 'video') {
        remoteVideo.srcObject = event.streams[0];
    } else if (event.track.kind === 'audio') {
        const audioElement = document.createElement('audio');
        audioElement.srcObject = event.streams[0];
        audioElement.autoplay = true;
        document.body.appendChild(audioElement);
    }
};

// Handling WebSocket messages
ws.onmessage = async (event) => {
    const message = JSON.parse(event.data);

    switch (message.type) {
        case 'offer':
            await pc.setRemoteDescription(new RTCSessionDescription(message.payload));
            const answer = await pc.createAnswer();
            await pc.setLocalDescription(answer);
            ws.send(JSON.stringify({ type: 'answer', payload: answer }));
            break;

        case 'answer':
            await pc.setRemoteDescription(new RTCSessionDescription(message.payload));
            break;

        case 'candidate':
            await pc.addIceCandidate(new RTCIceCandidate(message.payload));
            break;
    }
};

// Send ICE candidates to the server
pc.onicecandidate = (event) => {
    if (event.candidate) {
        ws.send(JSON.stringify({ type: 'candidate', payload: event.candidate }));
    }
};