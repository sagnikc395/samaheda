package main

type Signals int

const (
	EXIT_SUCCESS Signals = iota
	EXIT_FAILURE
)

const READLINE_BUFFERSIZE = 1024

const HOST_NAME_MAX = 512
const LOGIN_NMAE_MAX = 512
