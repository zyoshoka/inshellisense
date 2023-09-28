// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.
//
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package specs

import (
	"github.com/microsoft/clac/autocomplete/model"
)

func init() {
	Specs["ffmpeg"] = model.Subcommand{
		Name:        []string{"ffmpeg"},
		Description: `Play, record, convert, and stream audio and video`,
		Args: []model.Arg{{
			Templates:   []model.Template{model.TemplateFilepaths},
			Name:        "outfile",
			Description: `Output file`,
		}},
		Options: []model.Option{{
			Name:        []string{"-i"},
			Description: `Input file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "infile",
			}},
		}, {
			Name:        []string{"-L"},
			Description: `Show license`,
		}, {
			Name:        []string{"-h"},
			Description: `Show help`,
			Args: []model.Arg{{
				Name: "topic",
			}},
		}, {
			Name:        []string{"-?"},
			Description: `Show help`,
			Args: []model.Arg{{
				Name: "topic",
			}},
		}, {
			Name:        []string{"-help"},
			Description: `Show help`,
			Args: []model.Arg{{
				Name: "topic",
			}},
		}, {
			Name:        []string{"--help"},
			Description: `Show help`,
			Args: []model.Arg{{
				Name: "topic",
			}},
		}, {
			Name:        []string{"-version"},
			Description: `Show version`,
		}, {
			Name:        []string{"-buildconf"},
			Description: `Show build configuration`,
		}, {
			Name:        []string{"-formats"},
			Description: `Show available formats`,
		}, {
			Name:        []string{"-muxers"},
			Description: `Show available muxers`,
		}, {
			Name:        []string{"-demuxers"},
			Description: `Show available demuxers`,
		}, {
			Name:        []string{"-devices"},
			Description: `Show available devices`,
		}, {
			Name:        []string{"-codecs"},
			Description: `Show available codecs`,
		}, {
			Name:        []string{"-decoders"},
			Description: `Show available decoders`,
		}, {
			Name:        []string{"-encoders"},
			Description: `Show available encoders`,
		}, {
			Name:        []string{"-bsfs"},
			Description: `Show available bit stream filters`,
		}, {
			Name:        []string{"-protocols"},
			Description: `Show available protocols`,
		}, {
			Name:        []string{"-filters"},
			Description: `Show available filters`,
		}, {
			Name:        []string{"-pix_fmts"},
			Description: `Show available pixel formats`,
		}, {
			Name:        []string{"-layouts"},
			Description: `Show standard channel layouts`,
		}, {
			Name:        []string{"-sample_fmts"},
			Description: `Show available audio sample formats`,
		}, {
			Name:        []string{"-dispositions"},
			Description: `Show available stream dispositions`,
		}, {
			Name:        []string{"-colors"},
			Description: `Show available color names`,
		}, {
			Name:        []string{"-sources"},
			Description: `List sources of the input device`,
			Args: []model.Arg{{
				Name:      "device",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-sinks"},
			Description: `List sinks of the output device`,
			Args: []model.Arg{{
				Name:      "device",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-hwaccels"},
			Description: `Show available HW acceleration methods`,
		}, {
			Name:        []string{"-loglevel"},
			Description: `Set logging level`,
			Args: []model.Arg{{
				Name: "loglevel",
			}},
		}, {
			Name:        []string{"-v"},
			Description: `Set logging level`,
			Args: []model.Arg{{
				Name: "loglevel",
			}},
		}, {
			Name:        []string{"-report"},
			Description: `Generate a report`,
		}, {
			Name:        []string{"-max_alloc"},
			Description: `Set maximum size of a single allocated block`,
			Args: []model.Arg{{
				Name: "bytes",
			}},
		}, {
			Name:        []string{"-y"},
			Description: `Overwrite output files`,
		}, {
			Name:        []string{"-n"},
			Description: `Never overwrite output files`,
		}, {
			Name:        []string{"-ignore_unknown"},
			Description: `Ignore unknown stream types`,
		}, {
			Name:        []string{"-filter_threads"},
			Description: `Number of non-complex filter threads`,
		}, {
			Name:        []string{"-filter_complex_threads"},
			Description: `Number of threads for -filter_complex`,
		}, {
			Name:        []string{"-stats"},
			Description: `Print progress report during encoding`,
		}, {
			Name:        []string{"-max_error_rate"},
			Description: `Ratio of decoding errors (0.0: no errors, 1.0: 100% errors) above which ffmpeg returns an error instead of success`,
			Args: []model.Arg{{
				Name: "maximum error rate",
			}},
		}, {
			Name:        []string{"-vol"},
			Description: `Change audio volume (256=normal)`,
			Args: []model.Arg{{
				Name: "volume",
			}},
		}, {
			Name:        []string{"-cpuflags"},
			Description: `Force specific cpu flags`,
			Args: []model.Arg{{
				Name: "flags",
			}},
		}, {
			Name:        []string{"-cpucount"},
			Description: `Force specific cpu count`,
			Args: []model.Arg{{
				Name: "count",
			}},
		}, {
			Name:        []string{"-hide_banner"},
			Description: `Do not show program banner`,
			Args: []model.Arg{{
				Name: "hide_banner",
			}},
		}, {
			Name:        []string{"-copy_unknown"},
			Description: `Copy unknown stream types`,
		}, {
			Name:        []string{"-recast_media"},
			Description: `Allow recasting stream type in order to force a decoder of different media type`,
		}, {
			Name:        []string{"-benchmark"},
			Description: `Add timings for benchmarking`,
		}, {
			Name:        []string{"-benchmark_all"},
			Description: `Add timings for each task`,
		}, {
			Name:        []string{"-progress"},
			Description: `Write program-readable progress information`,
			Args: []model.Arg{{
				Name: "url",
			}},
		}, {
			Name:        []string{"-stdin"},
			Description: `Enable or disable interaction on standard input`,
		}, {
			Name:        []string{"-timelimit"},
			Description: `Set max runtime in seconds in CPU user time`,
			Args: []model.Arg{{
				Name: "limit",
			}},
		}, {
			Name:        []string{"-dump"},
			Description: `Dump each input packet`,
		}, {
			Name:        []string{"-hex"},
			Description: `When dumping packets, also dump the payload`,
		}, {
			Name:        []string{"-vsync"},
			Description: `Set video sync method globally; deprecated, use -fps_mode`,
		}, {
			Name:        []string{"-frame_drop_threshold"},
			Description: `Frame drop threshold`,
		}, {
			Name:        []string{"-async"},
			Description: `Audio sync method`,
		}, {
			Name:        []string{"-adrift_threshold"},
			Description: `Audio drift threshold`,
			Args: []model.Arg{{
				Name: "threshold",
			}},
		}, {
			Name:        []string{"-copyts"},
			Description: `Copy timestamps`,
		}, {
			Name:        []string{"-start_at_zero"},
			Description: `Shift input timestamps to start at 0 when using copyts`,
		}, {
			Name:        []string{"-copytb"},
			Description: `Copy input stream time base when stream copying`,
			Args: []model.Arg{{
				Name: "mode",
			}},
		}, {
			Name:        []string{"-dts_delta_threshold"},
			Description: `Timestamp discontinuity delta threshold`,
			Args: []model.Arg{{
				Name: "threshold",
			}},
		}, {
			Name:        []string{"-dts_error_threshold"},
			Description: `Timestamp error delta threshold`,
			Args: []model.Arg{{
				Name: "threshold",
			}},
		}, {
			Name:        []string{"-xerror"},
			Description: `Exit on error`,
			Args: []model.Arg{{
				Name: "error",
			}},
		}, {
			Name:        []string{"-abort_on"},
			Description: `Abort on the specified condition flags`,
			Args: []model.Arg{{
				Name: "flags",
			}},
		}, {
			Name:        []string{"-filter_complex"},
			Description: `Create a complex filtergraph`,
			Args: []model.Arg{{
				Name: "graph_description",
			}},
		}, {
			Name:        []string{"-lavfi"},
			Description: `Create a complex filtergraph`,
			Args: []model.Arg{{
				Name: "graph_description",
			}},
		}, {
			Name:        []string{"-filter_complex_script"},
			Description: `Read complex filtergraph description from a file`,
			Args: []model.Arg{{
				Name: "filename",
			}},
		}, {
			Name:        []string{"-auto_conversion_filters"},
			Description: `Enable automatic conversion filters globally`,
		}, {
			Name:        []string{"-stats_period"},
			Description: `Set the period at which ffmpeg updates stats and -progress output`,
			Args: []model.Arg{{
				Name: "time",
			}},
		}, {
			Name:        []string{"-debug_ts"},
			Description: `Print timestamp debugging info`,
		}, {
			Name:        []string{"-psnr"},
			Description: `Calculate PSNR of compressed frames`,
		}, {
			Name:        []string{"-vstats"},
			Description: `Dump video coding statistics to file`,
		}, {
			Name:        []string{"-vstats_file"},
			Description: `Dump video coding statistics to file`,
			Args: []model.Arg{{
				Name: "file",
			}},
		}, {
			Name:        []string{"-vstats_version"},
			Description: `Version of the vstats format to use`,
		}, {
			Name:        []string{"-qphist"},
			Description: `Show QP histogram`,
		}, {
			Name:        []string{"-sdp_file"},
			Description: `Specify a file in which to print sdp information`,
			Args: []model.Arg{{
				Name: "file",
			}},
		}, {
			Name:        []string{"-init_hw_device"},
			Description: `Initialise hardware device`,
			Args: []model.Arg{{
				Name: "args",
			}},
		}, {
			Name:        []string{"-filter_hw_device"},
			Description: `Set hardware device used when filtering`,
			Args: []model.Arg{{
				Name:      "device",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-f"},
			Description: `Force format`,
			Args: []model.Arg{{
				Name: "fmt",
			}},
		}, {
			Name:        []string{"-c"},
			Description: `Codec name`,
			Args: []model.Arg{{
				Name:      "codec",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-codec"},
			Description: `Codec name`,
			Args: []model.Arg{{
				Name:      "codec",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-pre"},
			Description: `Preset name`,
			Args: []model.Arg{{
				Name: "preset",
			}},
		}, {
			Name:        []string{"-map_metadata"},
			Description: `Set metadata information of outfile from infile`,
			Args: []model.Arg{{
				Name: "outfile[,metadata]:infile[,metadata]",
			}},
		}, {
			Name:        []string{"-t"},
			Description: `Record or transcode "duration" seconds of audio/video`,
			Args: []model.Arg{{
				Name: "duration",
			}},
		}, {
			Name:        []string{"-to"},
			Description: `Record or transcode stop time`,
			Args: []model.Arg{{
				Name: "time_stop",
			}},
		}, {
			Name:        []string{"-fs"},
			Description: `Set the limit file size in bytes`,
			Args: []model.Arg{{
				Name: "limit_size",
			}},
		}, {
			Name:        []string{"-ss"},
			Description: `Set the start time offset`,
			Args: []model.Arg{{
				Name: "time_off",
			}},
		}, {
			Name:        []string{"-sseof"},
			Description: `Set the start time offset relative to EOF`,
			Args: []model.Arg{{
				Name: "time_off",
			}},
		}, {
			Name:        []string{"-seek_timestamp"},
			Description: `Enable/disable seeking by timestamp with -ss`,
		}, {
			Name:        []string{"-timestamp"},
			Description: `Set the recording timestamp ('now' to set the current time)`,
			Args: []model.Arg{{
				Name: "time",
			}},
		}, {
			Name:        []string{"-metadata"},
			Description: `Add metadata`,
			Args: []model.Arg{{
				Name: "string=string",
			}},
		}, {
			Name:        []string{"-program"},
			Description: `Add program with specified streams`,
			Args: []model.Arg{{
				Name: "title=string:st=number...",
			}},
		}, {
			Name:        []string{"-target"},
			Description: `Specify target file type ("vcd", "svcd", "dvd", "dv" or "dv50" with optional prefixes "pal-", "ntsc-" or "film-")`,
			Args: []model.Arg{{
				Name: "type",
			}},
		}, {
			Name:        []string{"-apad"},
			Description: `Audio pad`,
		}, {
			Name:        []string{"-frames"},
			Description: `Set the number of frames to output`,
			Args: []model.Arg{{
				Name: "number",
			}},
		}, {
			Name:        []string{"-filter"},
			Description: `Set stream filtergraph`,
			Args: []model.Arg{{
				Name: "filter_graph",
			}},
		}, {
			Name:        []string{"-filter_script"},
			Description: `Read stream filtergraph description from a file`,
			Args: []model.Arg{{
				Name: "filename",
			}},
		}, {
			Name:        []string{"-reinit_filter"},
			Description: `Reinit filtergraph on input parameter changes`,
		}, {
			Name:        []string{"-discard"},
			Description: `Discard`,
		}, {
			Name:        []string{"-disposition"},
			Description: `Disposition`,
		}, {
			Name:        []string{"-map"},
			Description: `Set input stream mapping`,
			Args: []model.Arg{{
				Name: "[-]input_file_id[:stream_specifier][,sync_file_id[:stream_specifier]]",
			}},
		}, {
			Name:        []string{"-map_channel"},
			Description: `Map an audio channel from one stream to another`,
			Args: []model.Arg{{
				Name: "file.stream.channel[:syncfile.syncstream]",
			}},
		}, {
			Name:        []string{"-map_chapters"},
			Description: `Set chapters mapping`,
			Args: []model.Arg{{
				Name: "input_file_index",
			}},
		}, {
			Name:        []string{"-accurate_seek"},
			Description: `Enable/disable accurate seeking with -ss`,
		}, {
			Name:        []string{"-isync"},
			Description: `Indicate the input index for sync reference`,
			Args: []model.Arg{{
				Name: "sync ref",
			}},
		}, {
			Name:        []string{"-itsoffset"},
			Description: `Set the input ts offset`,
			Args: []model.Arg{{
				Name: "time_off",
			}},
		}, {
			Name:        []string{"-itsscale"},
			Description: `Set the input ts scale`,
			Args: []model.Arg{{
				Name: "scale",
			}},
		}, {
			Name:        []string{"-dframes"},
			Description: `Set the number of data frames to output`,
			Args: []model.Arg{{
				Name: "number",
			}},
		}, {
			Name:        []string{"-re"},
			Description: `Read input at native frame rate; equivalent to -readrate 1`,
		}, {
			Name:        []string{"-readrate"},
			Description: `Read input at specified rate`,
			Args: []model.Arg{{
				Name: "speed",
			}},
		}, {
			Name:        []string{"-shortest"},
			Description: `Finish encoding within shortest input`,
		}, {
			Name:        []string{"-bitexact"},
			Description: `Bitexact mode`,
		}, {
			Name:        []string{"-copyinkf"},
			Description: `Copy initial non-keyframes`,
		}, {
			Name:        []string{"-copypriorss"},
			Description: `Copy or discard frames before start time`,
		}, {
			Name:        []string{"-tag"},
			Description: `Force codec tag/fourcc`,
			Args: []model.Arg{{
				Name: "fourcc/tag",
			}},
		}, {
			Name:        []string{"-q"},
			Description: `Use fixed quality scale (VBR)`,
			Args: []model.Arg{{
				Name: "q",
			}},
		}, {
			Name:        []string{"-qscale"},
			Description: `Use fixed quality scale (VBR)`,
			Args: []model.Arg{{
				Name: "q",
			}},
		}, {
			Name:        []string{"-profile"},
			Description: `Set profile`,
			Args: []model.Arg{{
				Name: "profile",
			}},
		}, {
			Name:        []string{"-attach"},
			Description: `Add an attachment to the output file`,
			Args: []model.Arg{{
				Name: "filename",
			}},
		}, {
			Name:        []string{"-dump_attachment"},
			Description: `Extract an attachment into a file`,
			Args: []model.Arg{{
				Name: "filename",
			}},
		}, {
			Name:        []string{"-stream_loop"},
			Description: `Set number of times input stream shall be looped`,
			Args: []model.Arg{{
				Name: "loop count",
			}},
		}, {
			Name:        []string{"-thread_queue_size"},
			Description: `Set the maximum number of queued packets from the demuxer`,
		}, {
			Name:        []string{"-find_stream_info"},
			Description: `Read and decode the streams to fill missing information with heuristics`,
		}, {
			Name:        []string{"-bits_per_raw_sample"},
			Description: `Set the number of bits per raw sample`,
			Args: []model.Arg{{
				Name: "number",
			}},
		}, {
			Name:        []string{"-autorotate"},
			Description: `Automatically insert correct rotate filters`,
		}, {
			Name:        []string{"-autoscale"},
			Description: `Automatically insert a scale filter at the end of the filter graph`,
		}, {
			Name:        []string{"-muxdelay"},
			Description: `Set the maximum demux-decode delay`,
			Args: []model.Arg{{
				Name: "seconds",
			}},
		}, {
			Name:        []string{"-muxpreload"},
			Description: `Set the initial demux-decode delay`,
			Args: []model.Arg{{
				Name: "seconds",
			}},
		}, {
			Name:        []string{"-time_base"},
			Description: `Set the desired time base hint for output stream (1:24, 1:48000 or 0.04166, 2.0833e-5)`,
			Args: []model.Arg{{
				Name: "ratio",
			}},
		}, {
			Name:        []string{"-enc_time_base"},
			Description: `Set the desired time base for the encoder (1:24, 1:48000 or 0.04166, 2.0833e-5). two special values are defined - 0 = use frame rate (video) or sample rate (audio),-1 = match source time base`,
			Args: []model.Arg{{
				Name: "ratio",
			}},
		}, {
			Name:        []string{"-bsf"},
			Description: `A comma-separated list of bitstream filters`,
			Args: []model.Arg{{
				Name: "bitstream_filters",
			}},
		}, {
			Name:        []string{"-fpre"},
			Description: `Set options from indicated preset file`,
			Args: []model.Arg{{
				Name: "filename",
			}},
		}, {
			Name:        []string{"-max_muxing_queue_size"},
			Description: `Maximum number of packets that can be buffered while waiting for all streams to initialize`,
			Args: []model.Arg{{
				Name: "packets",
			}},
		}, {
			Name:        []string{"-muxing_queue_data_threshold"},
			Description: `Set the threshold after which max_muxing_queue_size is taken into account`,
			Args: []model.Arg{{
				Name: "bytes",
			}},
		}, {
			Name:        []string{"-dcodec"},
			Description: `Force data codec ('copy' to copy stream)`,
			Args: []model.Arg{{
				Name:      "codec",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-vframes"},
			Description: `Set the number of video frames to output`,
			Args: []model.Arg{{
				Name: "number",
			}},
		}, {
			Name:        []string{"-r"},
			Description: `Set frame rate (Hz value, fraction or abbreviation)`,
			Args: []model.Arg{{
				Name: "rate",
			}},
		}, {
			Name:        []string{"-fpsmax"},
			Description: `Set max frame rate (Hz value, fraction or abbreviation)`,
			Args: []model.Arg{{
				Name: "rate",
			}},
		}, {
			Name:        []string{"-s"},
			Description: `Set frame size (WxH or abbreviation)`,
			Args: []model.Arg{{
				Name: "size",
			}},
		}, {
			Name:        []string{"-aspect"},
			Description: `Set aspect ratio (4:3, 16:9 or 1.3333, 1.7777)`,
			Args: []model.Arg{{
				Name: "aspect",
			}},
		}, {
			Name:        []string{"-vn"},
			Description: `Disable video`,
		}, {
			Name:        []string{"-vcodec"},
			Description: `Force video codec ('copy' to copy stream)`,
			Args: []model.Arg{{
				Name:      "codec",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-timecode"},
			Description: `Set initial TimeCode value`,
			Args: []model.Arg{{
				Name: "hh:mm:ss[:;.]ff",
			}},
		}, {
			Name:        []string{"-pass"},
			Description: `Select the pass number (1 to 3)`,
			Args: []model.Arg{{
				Name: "n",
			}},
		}, {
			Name:        []string{"-vf"},
			Description: `Set video filters`,
			Args: []model.Arg{{
				Name: "filter_graph",
			}},
		}, {
			Name:        []string{"-ab"},
			Description: `Audio bitrate (please use -b:a)`,
			Args: []model.Arg{{
				Name: "bitrate",
			}},
		}, {
			Name:        []string{"-b"},
			Description: `Video bitrate (please use -b:v)`,
			Args: []model.Arg{{
				Name: "bitrate",
			}},
		}, {
			Name:        []string{"-dn"},
			Description: `Disable data`,
		}, {
			Name:        []string{"-pix_fmt"},
			Description: `Set pixel format`,
			Args: []model.Arg{{
				Name: "format",
			}},
		}, {
			Name:        []string{"-rc_override"},
			Description: `Rate control override for specific intervals`,
			Args: []model.Arg{{
				Name: "override",
			}},
		}, {
			Name:        []string{"-passlogfile"},
			Description: `Select two pass log file name prefix`,
			Args: []model.Arg{{
				Name: "prefix",
			}},
		}, {
			Name:        []string{"-intra_matrix"},
			Description: `Specify intra matrix coeffs`,
			Args: []model.Arg{{
				Name: "matrix",
			}},
		}, {
			Name:        []string{"-inter_matrix"},
			Description: `Specify inter matrix coeffs`,
			Args: []model.Arg{{
				Name: "matrix",
			}},
		}, {
			Name:        []string{"-chroma_intra_matrix"},
			Description: `Specify intra matrix coeffs`,
			Args: []model.Arg{{
				Name: "matrix",
			}},
		}, {
			Name:        []string{"-top"},
			Description: `Top=1/bottom=0/auto=-1 field first`,
		}, {
			Name:        []string{"-vtag"},
			Description: `Force video tag/fourcc`,
			Args: []model.Arg{{
				Name: "fourcc/tag",
			}},
		}, {
			Name:        []string{"-fps_mode"},
			Description: `Set framerate mode for matching video streams; overrides vsync`,
		}, {
			Name:        []string{"-force_fps"},
			Description: `Force the selected framerate, disable the best supported framerate selection`,
		}, {
			Name:        []string{"-streamid"},
			Description: `Set the value of an outfile streamid`,
			Args: []model.Arg{{
				Name: "streamIndex:value",
			}},
		}, {
			Name:        []string{"-force_key_frames"},
			Description: `Force key frames at specified timestamps`,
			Args: []model.Arg{{
				Name: "timestamps",
			}},
		}, {
			Name:        []string{"-hwaccel"},
			Description: `Use HW accelerated decoding`,
			Args: []model.Arg{{
				Name: "hwaccel name",
			}},
		}, {
			Name:        []string{"-hwaccel_device"},
			Description: `Select a device for HW acceleration`,
			Args: []model.Arg{{
				Name: "devicename",
			}},
		}, {
			Name:        []string{"-hwaccel_output_format"},
			Description: `Select output format used with HW accelerated decoding`,
			Args: []model.Arg{{
				Name: "format",
			}},
		}, {
			Name:        []string{"-vbsf"},
			Description: `Deprecated`,
			Args: []model.Arg{{
				Name: "video bitstream_filters",
			}},
		}, {
			Name:        []string{"-vpre"},
			Description: `Set the video options to the indicated preset`,
			Args: []model.Arg{{
				Name: "preset",
			}},
		}, {
			Name:        []string{"-aframes"},
			Description: `Set the number of audio frames to output`,
			Args: []model.Arg{{
				Name: "number",
			}},
		}, {
			Name:        []string{"-aq"},
			Description: `Set audio quality (codec-specific)`,
			Args: []model.Arg{{
				Name: "quality",
			}},
		}, {
			Name:        []string{"-ar"},
			Description: `Set audio sampling rate (in Hz)`,
			Args: []model.Arg{{
				Name: "rate",
			}},
		}, {
			Name:        []string{"-ac"},
			Description: `Set number of audio channels`,
			Args: []model.Arg{{
				Name: "channels",
			}},
		}, {
			Name:        []string{"-an"},
			Description: `Disable audio`,
		}, {
			Name:        []string{"-acodec"},
			Description: `Force audio codec ('copy' to copy stream)`,
			Args: []model.Arg{{
				Name:      "codec",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-af"},
			Description: `Set audio filters`,
			Args: []model.Arg{{
				Name: "filter_graph",
			}},
		}, {
			Name:        []string{"-atag"},
			Description: `Force audio tag/fourcc`,
			Args: []model.Arg{{
				Name: "fourcc/tag",
			}},
		}, {
			Name:        []string{"-sample_fmt"},
			Description: `Set sample format`,
			Args: []model.Arg{{
				Name: "format",
			}},
		}, {
			Name:        []string{"-channel_layout"},
			Description: `Set channel layout`,
			Args: []model.Arg{{
				Name: "layout",
			}},
		}, {
			Name:        []string{"-ch_layout"},
			Description: `Set channel layout`,
			Args: []model.Arg{{
				Name: "layout",
			}},
		}, {
			Name:        []string{"-guess_layout_max"},
			Description: `Set the maximum number of channels to try to guess the channel layout`,
		}, {
			Name:        []string{"-absf"},
			Description: `Deprecated`,
			Args: []model.Arg{{
				Name: "audio bitstream_filters",
			}},
		}, {
			Name:        []string{"-apre"},
			Description: `Set the audio options to the indicated preset`,
			Args: []model.Arg{{
				Name: "preset",
			}},
		}, {
			Name:        []string{"-sn"},
			Description: `Disable subtitle`,
		}, {
			Name:        []string{"-scodec"},
			Description: `Force subtitle codec ('copy' to copy stream)`,
			Args: []model.Arg{{
				Name:      "codec",
				Generator: nil, // TODO: port over generator
			}},
		}, {
			Name:        []string{"-stag"},
			Description: `Force subtitle tag/fourcc`,
			Args: []model.Arg{{
				Name: "fourcc/tag",
			}},
		}, {
			Name:        []string{"-fix_sub_duration"},
			Description: `Fix subtitles duration`,
		}, {
			Name:        []string{"-canvas_size"},
			Description: `Set canvas size (WxH or abbreviation)`,
			Args: []model.Arg{{
				Name: "size",
			}},
		}, {
			Name:        []string{"-spre"},
			Description: `Set the subtitle options to the indicated preset`,
			Args: []model.Arg{{
				Name: "preset",
			}},
		}},
	}
}
