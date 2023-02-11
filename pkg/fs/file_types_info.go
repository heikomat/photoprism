package fs

// TypeInfo contains human-readable descriptions for supported file formats
var TypeInfo = map[Type]string{
	ImageRaw:        "Unprocessed Sensor Data",
	ImageDNG:        "Adobe Digital Negative",
	ImageJPEG:       "Joint Photographic Experts Group (JPEG)",
	ImagePNG:        "Portable Network Graphics",
	ImageGIF:        "Graphics Interchange Format",
	ImageTIFF:       "Tag Image File Format",
	ImageBMP:        "Bitmap",
	ImageMPO:        "Stereoscopic JPEG (3D)",
	ImageAVIF:       "AV1 Image File Format",
	ImageHEIF:       "High Efficiency Image File Format",
	ImageHEIC:       "High Efficiency Image Container",
	ImageWebP:       "Google WebP",
	VideoWebM:       "Google WebM",
	VideoMP2:        "MPEG 2 (H.262)",
	VideoAVC:        "Advanced Video Coding (H.264, MPEG-4 Part 10)",
	VideoHEVC:       "High Efficiency Video Coding (H.265)",
	VideoVVC:        "Versatile Video Coding (H.266)",
	VideoAV1:        "AOMedia Video 1",
	VideoMOV:        "Apple QuickTime",
	VideoMP4:        "Multimedia Container (MPEG-4 Part 14)",
	VideoAVI:        "Microsoft Audio Video Interleave",
	VideoASF:        "Advanced Systems Format ",
	VideoWMV:        "Windows Media",
	Video3GP:        "Mobile Multimedia Container (3G)",
	Video3G2:        "Mobile Multimedia Container (CDMA2000)",
	VideoFlash:      "Adobe Flash",
	VideoMKV:        "Matroska Multimedia Container",
	VideoMPG:        "Moving Picture Experts Group (MPEG)",
	VideoMJPG:       "Motion JPEG",
	VideoAVCHD:      "Advanced Video Coding High Definition (AVCHD)",
	VideoBDAV:       "Blu-ray MPEG-2 Transport Stream",
	VideoOGV:        "Ogg Media (OGG)",
	VectorSVG:       "Scalable Vector Graphics",
	VectorPS:        "Adobe PostScript",
	VectorEPS:       "Adobe Encapsulated PostScript",
	SidecarXMP:      "Adobe Extensible Metadata Platform",
	SidecarAAE:      "Apple Image Edits XML",
	SidecarXML:      "Extensible Markup Language",
	SidecarJSON:     "Serialized JSON Data (Exiftool, Google Photos)",
	SidecarYAML:     "Serialized YAML Data (Config, Metadata)",
	SidecarText:     "Plain Text",
	SidecarMarkdown: "Markdown Formatted Text",
	UnknownType:     "Other",
}
