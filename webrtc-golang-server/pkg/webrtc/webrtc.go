package webrtc

import (
	"log"

	"github.com/pion/webrtc/v3"
)

var peerConnection *webrtc.PeerConnection

func Init() {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
			// adding TURN server credentials here ---- for production
		},
	}

	var err error
	peerConnection, err = webrtc.NewPeerConnection(config)
	if err != nil {
		log.Fatal("Failed to create PeerConnection:", err)
	}

	peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			log.Println("New ICE candidate:", candidate)
			// Sending candidate to remote peer via signaling
		}
	})

	peerConnection.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		log.Println("Connection state changed:", state)
	})

	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		log.Println("Track received:", track.Kind())
<<<<<<< HEAD

		// Handle the track (e.g., play audio/video)
		// if track.Kind() == webrtc.RTPCodecTypeVideo {
		// 	go handleVideoTrack(track)
		// } else if track.Kind() == webrtc.RTPCodecTypeAudio {
		// 	go handleAudioTrack(track)
		// }
=======
		// track handling e.g. play audio/video
>>>>>>> df38c5970cdaadfe25d8e9bde799d55c18cc50d0
	})
}

// func handleVideoTrack(track *webrtc.TrackRemote) {
// 	for {
// 		rtpPacket, _, err := track.ReadRTP()
// 		if err != nil {
// 			log.Println("Error reading RTP packet:", err)
// 			return
// 		}

// 		// Send the RTP packet to the frontend for rendering
// 		// it requires integration with the frontend (e.g., using WebSockets or WebRTC data channels)
// 	}
// }

// func handleAudioTrack(track *webrtc.TrackRemote) {
// 	for {
// 		rtpPacket, _, err := track.ReadRTP()
// 		if err != nil {
// 			log.Println("Error reading RTP packet:", err)
// 			return
// 		}

//			// Send the RTP packet to the frontend for rendering
//		}
//	}
func CreateOffer() (*webrtc.SessionDescription, error) {
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		return nil, err
	}

	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		return nil, err
	}

	return &offer, nil
}

func SetRemoteDescription(desc webrtc.SessionDescription) error {
	return peerConnection.SetRemoteDescription(desc)
}

func AddICECandidate(candidate webrtc.ICECandidateInit) error {
	return peerConnection.AddICECandidate(candidate)
}
