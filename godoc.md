package pipe // import "github.com/GoLangsam/pipe"


FUNCTIONS

func AnyChan(inp ...Any) (out <-chan Any)
    AnyChan returns a channel to receive all inputs before close.

func AnyChanFuncErr(gen func() (Any, error)) (out <-chan Any)
    AnyChanFuncErr returns a channel to receive all results of generator `gen`
    until `err != nil` before close.

func AnyChanFuncNok(gen func() (Any, bool)) (out <-chan Any)
    AnyChanFuncNok returns a channel to receive all results of generator `gen`
    until `!ok` before close.

func AnyChanSlice(inp ...[]Any) (out <-chan Any)
    AnyChanSlice returns a channel to receive all inputs before close.

func AnyDaisyChaiN(
	inp chan Any,
	somany int,
	procs ...func(out chan<- Any, inp <-chan Any),
) (
	out chan Any,
)
    AnyDaisyChaiN returns a channel to receive all inp after having passed
    `somany` times thru the process(es) (`from` right `into` left) before close.

    Note: If `somany` is less than 1 or no `tubes` are provided, `out` shall
    receive elements from `inp` unaltered (as a convenience), thus making null
    values useful.

    Note: AnyDaisyChaiN(inp, 1, procs) <==> AnyDaisyChain(inp, procs)

func AnyDaisyChain(
	inp chan Any,
	procs ...func(out chan<- Any, inp <-chan Any),
) (
	out chan Any,
)
    AnyDaisyChain returns a channel to receive all inp after having passed thru
    the process(es) (`from` right `into` left) before close.

    Note: If no `tubes` are provided, `out` shall receive elements from `inp`
    unaltered (as a convenience), thus making a null value useful.

func AnyDone(inp <-chan Any, ops ...func(a Any)) (done <-chan struct{})
    AnyDone will apply every `op` to every `inp` and returns a channel to
    receive one signal upon close.

func AnyDoneFreq(inp <-chan Any) (freq <-chan map[Any]int64)
    AnyDoneFreq returns a channel to receive a frequency histogram (as a
    `map[Any]int64`) upon close.

func AnyDoneFreqAttr(inp <-chan Any, attr func(a Any) interface{}) (freq <-chan map[interface{}]int64)
    AnyDoneFreqAttr returns a channel to receive a frequency histogram (as a
    `map[interface{}]int64`) upon close.

    `attr` provides the key to the frequency map. If `nil` is passed as `attr`
    then Any is used as key.

func AnyDoneFunc(inp <-chan Any, acts ...func(a Any) Any) (done <-chan struct{})
    AnyDoneFunc will chain every `act` to every `inp` and returns a channel to
    receive one signal upon close.

func AnyDoneLeave(inp <-chan Any, wg AnyWaiter) (done <-chan struct{})
    AnyDoneLeave returns a channel to receive one signal after all throughput on
    `inp` has been registered as departure on the given `sync.WaitGroup` before
    close.

func AnyDoneSlice(inp <-chan Any) (done <-chan []Any)
    AnyDoneSlice returns a channel to receive a slice with every Any received on
    `inp` upon close.

    Note: Unlike AnyDone, AnyDoneSlice sends the fully accumulated slice, not
    just an event, once upon close of inp.

func AnyDoneWait(out chan<- Any, wg AnyWaiter) (done <-chan struct{})
    AnyDoneWait returns a channel to receive one signal after wg.Wait() has
    returned and out has been closed before close.

    Note: Use only *after* You've started flooding the facilities.

func AnyFan2(inp <-chan Any, inps ...Any) (out <-chan Any)
    AnyFan2 returns a channel to receive everything from `inp` as well as all
    inputs before close.

func AnyFan2Chan(inp <-chan Any, inp2 <-chan Any) (out <-chan Any)
    AnyFan2Chan returns a channel to receive everything from `inp` as well as
    everything from `inp2` before close. Note: AnyFan2Chan is nothing but
    AnyFanIn2

func AnyFan2FuncErr(inp <-chan Any, gen func() (Any, error)) (out <-chan Any)
    AnyFan2FuncErr returns a channel to receive everything from `inp` as well as
    all results of generator `gen` until `err != nil` before close.

func AnyFan2FuncNok(inp <-chan Any, gen func() (Any, bool)) (out <-chan Any)
    AnyFan2FuncNok returns a channel to receive everything from `inp` as well as
    all results of generator `gen` until `!ok` before close.

func AnyFan2Slice(inp <-chan Any, inps ...[]Any) (out <-chan Any)
    AnyFan2Slice returns a channel to receive everything from `inp` as well as
    all inputs before close.

func AnyFanIn(inps ...<-chan Any) (out <-chan Any)
    AnyFanIn returns a channel to receive all inputs arriving on variadic inps
    before close.

    Note: For each input one go routine is spawned to forward arrivals.

    See AnyFanIn1 in `fan-in1` for another implementation.

    Ref: https://blog.golang.org/pipelines
    Ref: https://github.com/QuentinPerez/go-stuff/channel/Fan-out-Fan-in/main.go

func AnyFanIn1(inpS ...<-chan Any) (out <-chan Any)
    AnyFanIn1 returns a channel to receive all inputs arriving on variadic inps
    before close.

    Note: Only one go routine is used for all receives,
    which keeps trying open inputs in round-robin fashion
    until all inputs are closed.

    See AnyFanIn in `fan-in` for another implementation.

func AnyFanIn2(inp, inp2 <-chan Any) (out <-chan Any)
    AnyFanIn2 returns a channel to receive all from both `inp` and `inp2` before
    close.

func AnyFanOut(inp <-chan Any, size int) (outS [](<-chan Any))
    AnyFanOut returns a slice (of size = size) of channels each of which shall
    receive any inp before close.

func AnyFini(ops ...func(a Any)) func(inp <-chan Any) (done <-chan struct{})
    AnyFini returns a closure around `AnyDone(_, ops...)`.

func AnyFiniFunc(acts ...func(a Any) Any) func(inp <-chan Any) (done <-chan struct{})
    AnyFiniFunc returns a closure around `AnyDoneFunc(_, acts...)`.

func AnyFiniLeave(wg AnyWaiter) func(inp <-chan Any) (done <-chan struct{})
    AnyFiniLeave returns a closure around `AnyDoneLeave(_, wg)` registering
    throughput as departure on the given `sync.WaitGroup`.

func AnyFiniSlice() func(inp <-chan Any) (done <-chan []Any)
    AnyFiniSlice returns a closure around `AnyDoneSlice(_)`.

func AnyFiniWait(wg AnyWaiter) func(out chan<- Any) (done <-chan struct{})
    AnyFiniWait returns a closure around `AnyDoneWait(_, wg)`.

func AnyFork(inp <-chan Any) (out1, out2 <-chan Any)
    AnyFork returns two channels either of which is to receive every result of
    inp before close.

func AnyForkSeen(inp <-chan Any) (new, old <-chan Any)
    AnyForkSeen returns two channels, `new` and `old`, where `new` is to receive
    all `inp` not been seen before and `old` all `inp` seen before (internally
    growing a `sync.Map` to discriminate) until close.

func AnyForkSeenAttr(inp <-chan Any, attr func(a Any) interface{}) (new, old <-chan Any)
    AnyForkSeenAttr returns two channels, `new` and `old`, where `new` is to
    receive all `inp` whose attribute `attr` has not been seen before and `old`
    all `inp` seen before (internally growing a `sync.Map` to discriminate)
    until close.

func AnyJoin(out chan<- Any, inp ...Any) (done <-chan struct{})
    AnyJoin sends inputs on the given out channel and returns a done channel to
    receive one signal when inp has been drained

func AnyJoinChan(out chan<- Any, inp <-chan Any) (done <-chan struct{})
    AnyJoinChan sends inputs on the given out channel and returns a done channel
    to receive one signal when inp has been drained

func AnyJoinSlice(out chan<- Any, inp ...[]Any) (done <-chan struct{})
    AnyJoinSlice sends inputs on the given out channel and returns a done
    channel to receive one signal when inp has been drained

func AnyMakeChan() (out chan Any)
    AnyMakeChan returns a new open channel (simply a 'chan Any' that is).

    Note: No 'Any-producer' is launched here yet! (as is in all the other
    functions).

    This is useful to easily create corresponding variables such as:

    var myAnyPipelineStartsHere := AnyMakeChan() // ... lot's of code to design
    and build Your favourite "myAnyWorkflowPipeline"

    // ...
    // ... *before* You start pouring data into it, e.g. simply via:
    for drop := range water {

    myAnyPipelineStartsHere <- drop

    }

    close(myAnyPipelineStartsHere)

    Hint: especially helpful, if Your piping library operates on some hidden
    (non-exported) type (or on a type imported from elsewhere - and You don't
    want/need or should(!) have to care.)

    Note: as always (except for AnyPipeBuffer) the channel is unbuffered.

func AnyMerge(less func(i, j Any) bool, inps ...<-chan Any) (out <-chan Any)
    AnyMerge returns a channel to receive all inputs sorted and free of
    duplicates. Each input channel needs to be sorted ascending and free of
    duplicates. The passed binary boolean function `less` defines the applicable
    order.

    Note: If no inputs are given, a closed channel is returned.

func AnyPair(inp <-chan Any) (out1, out2 <-chan Any)
    AnyPair returns a pair of channels to receive every result of inp before
    close.

    Note: Yes, it is a VERY simple fanout - but sometimes all You need.

func AnyPipe(inp <-chan Any, ops ...func(a Any)) (out <-chan Any)
    AnyPipe will apply every `op` to every `inp` and returns a channel to
    receive each `inp` before close.

    Note: For functional people, this 'could' be named `AnyMap`. Just: 'map' has
    a very different meaning in go lang.

func AnyPipeAdjust(inp <-chan Any, sizes ...int) (out <-chan Any)
    AnyPipeAdjust returns a channel to receive all `inp` buffered by a
    AnySendProxy process before close.

func AnyPipeBuffered(inp <-chan Any, cap int) (out <-chan Any)
    AnyPipeBuffered returns a buffered channel with capacity `cap` to receive
    all `inp` before close.

func AnyPipeDone(inp <-chan Any) (out <-chan Any, done <-chan struct{})
    AnyPipeDone returns a channel to receive every `inp` before close and a
    channel to signal this closing.

func AnyPipeEnter(inp <-chan Any, wg AnyWaiter) (out <-chan Any)
    AnyPipeEnter returns a channel to receive all `inp` and registers throughput
    as arrival on the given `sync.WaitGroup` until close.

func AnyPipeFunc(inp <-chan Any, acts ...func(a Any) Any) (out <-chan Any)
    AnyPipeFunc will chain every `act` to every `inp` and returns a channel to
    receive each result before close.

func AnyPipeFuncMany(inp <-chan Any, act func(a Any) Any, many int) (out <-chan Any)
    AnyPipeFuncMany returns a channel to receive every result of action `act`
    applied to `inp` by `many` parallel processing goroutines before close.

    ref: database/sql/sql_test.go
    ref: cmd/compile/internal/gc/noder.go

func AnyPipeLeave(inp <-chan Any, wg AnyWaiter) (out <-chan Any)
    AnyPipeLeave returns a channel to receive all `inp` and registers throughput
    as departure on the given `sync.WaitGroup` until close.

func AnyPipeSeen(inp <-chan Any) (out <-chan Any)
    AnyPipeSeen returns a channel to receive all `inp` not been seen before
    while silently dropping everything seen before (internally growing a
    `sync.Map` to discriminate) until close. Note: AnyPipeFilterNotSeenYet might
    be a better name, but is fairly long.

func AnyPipeSeenAttr(inp <-chan Any, attr func(a Any) interface{}) (out <-chan Any)
    AnyPipeSeenAttr returns a channel to receive all `inp` whose attribute
    `attr` has not been seen before while silently dropping everything seen
    before (internally growing a `sync.Map` to discriminate) until close. Note:
    AnyPipeFilterAttrNotSeenYet might be a better name, but is fairly long.

func AnyPlug(inp <-chan Any, stop <-chan struct{}) (out <-chan Any, done <-chan struct{})
    AnyPlug returns a channel to receive every `inp` before close and a channel
    to signal this closing. Upon receipt of a stop signal, output is immediately
    closed, and for graceful termination any remaining input is drained before
    done is signalled.

func AnyPlugAfter(inp <-chan Any, after <-chan time.Time) (out <-chan Any, done <-chan struct{})
    AnyPlugAfter returns a channel to receive every `inp` before close and a
    channel to signal this closing. Upon receipt of a time signal (e.g. from
    `time.After(...)`), output is immediately closed, and for graceful
    termination any remaining input is drained before done is signalled.

func AnySame(same func(a, b Any) bool, inp, inp2 <-chan Any) (out <-chan bool)
    AnySame reads values from two channels in lockstep and iff they have the
    same contents then `true` is sent on the returned bool channel before close.

func AnySendProxy(out chan<- Any, sizes ...int) chan<- Any
    AnySendProxy returns a channel to serve as a sending proxy to 'out'. Uses a
    goroutine to receive values from 'out' and store them in an expanding
    buffer, so that sending to 'out' never blocks.

    Note: the expanding buffer is implemented via "container/ring"

    Note: AnySendProxy is kept for the Sieve example and other dynamic use to be
    discovered even so it does not fit the pipe tube pattern as AnyPipeAdjust
    does.

func AnyStrew(inp <-chan Any, size int) (outS [](<-chan Any))
    AnyStrew returns a slice (of size = size) of channels one of which shall
    receive each inp before close.

func AnyTube(ops ...func(a Any)) (tube func(inp <-chan Any) (out <-chan Any))
    AnyTube returns a closure around PipeAny (_, ops...).

func AnyTubeAdjust(sizes ...int) (tube func(inp <-chan Any) (out <-chan Any))
    AnyTubeAdjust returns a closure around AnyPipeAdjust (_, sizes ...int).

func AnyTubeBuffered(cap int) (tube func(inp <-chan Any) (out <-chan Any))
    AnyTubeBuffered returns a closure around PipeAnyBuffer (_, cap).

func AnyTubeEnter(wg AnyWaiter) (tube func(inp <-chan Any) (out <-chan Any))
    AnyTubeEnter returns a closure around AnyPipeEnter (_, wg) registering
    throughput as arrival on the given `sync.WaitGroup`.

func AnyTubeFunc(acts ...func(a Any) Any) (tube func(inp <-chan Any) (out <-chan Any))
    AnyTubeFunc returns a closure around PipeAnyFunc (_, acts...).

func AnyTubeLeave(wg AnyWaiter) (tube func(inp <-chan Any) (out <-chan Any))
    AnyTubeLeave returns a closure around AnyPipeLeave (_, wg) registering
    throughput as departure on the given `sync.WaitGroup`.

func AnyTubeSeen() (tube func(inp <-chan Any) (out <-chan Any))
    AnyTubeSeen returns a closure around AnyPipeSeen() (silently dropping every
    Any seen before).

func AnyTubeSeenAttr(attr func(a Any) interface{}) (tube func(inp <-chan Any) (out <-chan Any))
    AnyTubeSeenAttr returns a closure around AnyPipeSeenAttr() (silently
    dropping every Any whose attribute `attr` was seen before).


TYPES

type Any generic.Type
    Any is the generic type flowing thru the pipe network.

type AnyProc func(into chan<- Any, from <-chan Any)
    AnyProc is the signature of the inner process of any linear pipe-network

    Example: the identity proc:

    samesame := func(into chan<- Any, from <-chan Any) { into <- <-from } Note:
    type AnyProc is provided for documentation purpose only. The implementation
    uses the explicit function signature in order to avoid some genny-related
    issue.

    Note: In https://talks.golang.org/2012/waza.slide#40

    Rob Pike uses a AnyProc named `worker`.

type AnyWaiter interface {
	Add(delta int)
	Done()
	Wait()
}
    AnyWaiter - as implemented by `*sync.WaitGroup` - attends Flapdoors and
    keeps counting who enters and who leaves.

    Use AnyDoneWait to learn about when the facilities are closed.

    Note: You may also use Your provided `*sync.WaitGroup.Wait()` to know when
    to close the facilities. Just: AnyDoneWait is more convenient as it also
    closes the primary channel for You.

    Just make sure to have _all_ entrances and exits attended, and `Wait()` only
    *after* You've started flooding the facilities.

