// This file is automatically generated by qtc from "extra.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/extra.qtpl:1
package templates

//line templates/extra.qtpl:1
import (
	"github.com/knq/chromedp/cmd/chromedp-gen/internal"
)

// ExtraTimestampTemplate is a special template for the Timestamp type that
// defines its JSON unmarshaling.

//line templates/extra.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/extra.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/extra.qtpl:7
func StreamExtraTimestampTemplate(qw422016 *qt422016.Writer, t *internal.Type, d *internal.Domain) {
	//line templates/extra.qtpl:8
	typ := t.IDorName()
	bootstamp := t.TimestampType == internal.TimestampTypeBootstamp
	timeRes := "time.Millisecond"
	if t.TimestampType != internal.TimestampTypeMillisecond {
		timeRes = "time.Second"
	}

	//line templates/extra.qtpl:14
	qw422016.N().S(`
// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t `)
	//line templates/extra.qtpl:16
	qw422016.N().S(typ)
	//line templates/extra.qtpl:16
	qw422016.N().S(`) MarshalEasyJSON(out *jwriter.Writer) {
	v := `)
	//line templates/extra.qtpl:17
	if bootstamp {
		//line templates/extra.qtpl:17
		qw422016.N().S(`float64(time.Time(t).Sub(sysutil.BootTime()))/float64(time.Second)`)
		//line templates/extra.qtpl:17
	} else {
		//line templates/extra.qtpl:17
		qw422016.N().S(`float64(time.Time(t).UnixNano()/int64(`)
		//line templates/extra.qtpl:17
		qw422016.N().S(timeRes)
		//line templates/extra.qtpl:17
		qw422016.N().S(`))`)
		//line templates/extra.qtpl:17
	}
	//line templates/extra.qtpl:17
	qw422016.N().S(`

	out.Buffer.EnsureSpace(20)
	out.Buffer.Buf = strconv.AppendFloat(out.Buffer.Buf, v, 'f', -1, 64)
}

// MarshalJSON satisfies json.Marshaler.
func (t `)
	//line templates/extra.qtpl:24
	qw422016.N().S(typ)
	//line templates/extra.qtpl:24
	qw422016.N().S(`) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *`)
	//line templates/extra.qtpl:29
	qw422016.N().S(typ)
	//line templates/extra.qtpl:29
	qw422016.N().S(`) UnmarshalEasyJSON(in *jlexer.Lexer) {`)
	//line templates/extra.qtpl:29
	if bootstamp {
		//line templates/extra.qtpl:29
		qw422016.N().S(`
	*t = `)
		//line templates/extra.qtpl:30
		qw422016.N().S(typ)
		//line templates/extra.qtpl:30
		qw422016.N().S(`(sysutil.BootTime().Add(time.Duration(in.Float64()*float64(time.Second))))`)
		//line templates/extra.qtpl:30
	} else {
		//line templates/extra.qtpl:30
		qw422016.N().S(`
	*t = `)
		//line templates/extra.qtpl:31
		qw422016.N().S(typ)
		//line templates/extra.qtpl:31
		qw422016.N().S(`(time.Unix(0, int64(in.Float64()*float64(`)
		//line templates/extra.qtpl:31
		qw422016.N().S(timeRes)
		//line templates/extra.qtpl:31
		qw422016.N().S(`))))`)
		//line templates/extra.qtpl:31
	}
	//line templates/extra.qtpl:31
	qw422016.N().S(`
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *`)
	//line templates/extra.qtpl:35
	qw422016.N().S(typ)
	//line templates/extra.qtpl:35
	qw422016.N().S(`) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
`)
//line templates/extra.qtpl:38
}

//line templates/extra.qtpl:38
func WriteExtraTimestampTemplate(qq422016 qtio422016.Writer, t *internal.Type, d *internal.Domain) {
	//line templates/extra.qtpl:38
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/extra.qtpl:38
	StreamExtraTimestampTemplate(qw422016, t, d)
	//line templates/extra.qtpl:38
	qt422016.ReleaseWriter(qw422016)
//line templates/extra.qtpl:38
}

//line templates/extra.qtpl:38
func ExtraTimestampTemplate(t *internal.Type, d *internal.Domain) string {
	//line templates/extra.qtpl:38
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/extra.qtpl:38
	WriteExtraTimestampTemplate(qb422016, t, d)
	//line templates/extra.qtpl:38
	qs422016 := string(qb422016.B)
	//line templates/extra.qtpl:38
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/extra.qtpl:38
	return qs422016
//line templates/extra.qtpl:38
}

// ExtraFrameTemplate is a special template for the Page.Frame type, adding FrameState.

//line templates/extra.qtpl:41
func StreamExtraFrameTemplate(qw422016 *qt422016.Writer) {
	//line templates/extra.qtpl:41
	qw422016.N().S(`
// FrameState is the state of a Frame.
type FrameState uint16

// FrameState enum values.
const (
    FrameDOMContentEventFired FrameState = 1 << (15 - iota)
    FrameLoadEventFired
    FrameAttached
    FrameNavigated
    FrameLoading
    FrameScheduledNavigation
)

// frameStateNames are the names of the frame states.
var frameStateNames = map[FrameState]string{
    FrameDOMContentEventFired: "DOMContentEventFired",
    FrameLoadEventFired:       "LoadEventFired",
    FrameAttached:             "Attached",
    FrameNavigated:            "Navigated",
	FrameLoading:			   "Loading",
    FrameScheduledNavigation:  "ScheduledNavigation",
}

// String satisfies stringer interface.
func (fs FrameState) String() string {
    var s []string
    for k, v := range frameStateNames {
        if fs&k != 0 {
            s = append(s, v)
        }
    }
    return "[" + strings.Join(s, " ") + "]"
}

// EmptyFrameID is the "non-existent" frame id.
const EmptyFrameID = FrameID("")
`)
//line templates/extra.qtpl:78
}

//line templates/extra.qtpl:78
func WriteExtraFrameTemplate(qq422016 qtio422016.Writer) {
	//line templates/extra.qtpl:78
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/extra.qtpl:78
	StreamExtraFrameTemplate(qw422016)
	//line templates/extra.qtpl:78
	qt422016.ReleaseWriter(qw422016)
//line templates/extra.qtpl:78
}

//line templates/extra.qtpl:78
func ExtraFrameTemplate() string {
	//line templates/extra.qtpl:78
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/extra.qtpl:78
	WriteExtraFrameTemplate(qb422016)
	//line templates/extra.qtpl:78
	qs422016 := string(qb422016.B)
	//line templates/extra.qtpl:78
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/extra.qtpl:78
	return qs422016
//line templates/extra.qtpl:78
}

// ExtraNodeTemplate is a special template for the DOM.Node type, adding NodeState.

//line templates/extra.qtpl:81
func StreamExtraNodeTemplate(qw422016 *qt422016.Writer) {
	//line templates/extra.qtpl:81
	qw422016.N().S(`
// AttributeValue returns the named attribute for the node.
func (n *Node) AttributeValue(name string) string {
	n.RLock()
	defer n.RUnlock()

	for i := 0; i < len(n.Attributes); i+=2 {
		if n.Attributes[i] == name  {
			return n.Attributes[i+1]
		}
	}

	return ""
}

// xpath builds the xpath string.
func (n *Node) xpath(stopAtDocument, stopAtID bool) string {
	n.RLock()
	defer n.RUnlock()

	p := ""
	pos := ""
	id := n.AttributeValue("id")
	switch {
	case n.Parent == nil:
		return n.LocalName

	case stopAtDocument && n.NodeType == NodeTypeDocument:
		return ""

	case stopAtID && id != "":
		p = "/"
		pos = `)
	//line templates/extra.qtpl:81
	qw422016.N().S("`")
	//line templates/extra.qtpl:81
	qw422016.N().S(`[@id='`)
	//line templates/extra.qtpl:81
	qw422016.N().S("`")
	//line templates/extra.qtpl:81
	qw422016.N().S(`+id+`)
	//line templates/extra.qtpl:81
	qw422016.N().S("`")
	//line templates/extra.qtpl:81
	qw422016.N().S(`']`)
	//line templates/extra.qtpl:81
	qw422016.N().S("`")
	//line templates/extra.qtpl:81
	qw422016.N().S(`

	case n.Parent != nil:
		var i int
		var found bool

		n.Parent.RLock()
		for j := 0; j < len(n.Parent.Children); j++ {
			if n.Parent.Children[j].LocalName == n.LocalName {
				i++
			}
			if n.Parent.Children[j].NodeID == n.NodeID {
				found = true
				break
			}
		}
		n.Parent.RUnlock()

		if found {
			pos = "["+strconv.Itoa(i)+"]"
		}

		p = n.Parent.xpath(stopAtDocument, stopAtID)
	}

	return  p + "/" + n.LocalName + pos
}

// PartialXPathByID returns the partial XPath for the node, stopping at the
// first parent with an id attribute or at nearest parent document node.
func (n *Node) PartialXPathByID() string {
	return n.xpath(true, true)
}

// PartialXPath returns the partial XPath for the node, stopping at the nearest
// parent document node.
func (n *Node) PartialXPath() string {
	return n.xpath(true, false)
}

// FullXPathByID returns the full XPath for the node, stopping at the top most
// document root or at the closest parent node with an id attribute.
func (n *Node) FullXPathByID() string {
	return n.xpath(false, true)
}

// FullXPath returns the full XPath for the node, stopping only at the top most
// document root.
func (n *Node) FullXPath() string {
	return n.xpath(false, false)
}

// NodeState is the state of a DOM node.
type NodeState uint8

// NodeState enum values.
const (
    NodeReady NodeState = 1 << (7 - iota)
	NodeVisible
	NodeHighlighted
)

// nodeStateNames are the names of the node states.
var nodeStateNames = map[NodeState]string{
    NodeReady:		 "Ready",
    NodeVisible:     "Visible",
    NodeHighlighted: "Highlighted",
}

// String satisfies stringer interface.
func (ns NodeState) String() string {
    var s []string
    for k, v := range nodeStateNames {
        if ns&k != 0 {
            s = append(s, v)
        }
    }
    return "[" + strings.Join(s, " ") + "]"
}

// EmptyNodeID is the "non-existent" node id.
const EmptyNodeID = NodeID(0)
`)
//line templates/extra.qtpl:195
}

//line templates/extra.qtpl:195
func WriteExtraNodeTemplate(qq422016 qtio422016.Writer) {
	//line templates/extra.qtpl:195
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/extra.qtpl:195
	StreamExtraNodeTemplate(qw422016)
	//line templates/extra.qtpl:195
	qt422016.ReleaseWriter(qw422016)
//line templates/extra.qtpl:195
}

//line templates/extra.qtpl:195
func ExtraNodeTemplate() string {
	//line templates/extra.qtpl:195
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/extra.qtpl:195
	WriteExtraNodeTemplate(qb422016)
	//line templates/extra.qtpl:195
	qs422016 := string(qb422016.B)
	//line templates/extra.qtpl:195
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/extra.qtpl:195
	return qs422016
//line templates/extra.qtpl:195
}

// ExtraFixStringUnmarshaler is a template that forces values to be parsed properly.

//line templates/extra.qtpl:198
func StreamExtraFixStringUnmarshaler(qw422016 *qt422016.Writer, typ, parseFunc, extra string) {
	//line templates/extra.qtpl:198
	qw422016.N().S(`
// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *`)
	//line templates/extra.qtpl:200
	qw422016.N().S(typ)
	//line templates/extra.qtpl:200
	qw422016.N().S(`) UnmarshalEasyJSON(in *jlexer.Lexer) {
	buf := in.Raw()
	if l := len(buf); l > 2 && buf[0] == '"' && buf[l-1] == '"' {
		buf = buf[1:l-1]
	}
`)
	//line templates/extra.qtpl:205
	if parseFunc != "" {
		//line templates/extra.qtpl:205
		qw422016.N().S(`
	v, err := strconv.`)
		//line templates/extra.qtpl:206
		qw422016.N().S(parseFunc)
		//line templates/extra.qtpl:206
		qw422016.N().S(`(string(buf)`)
		//line templates/extra.qtpl:206
		qw422016.N().S(extra)
		//line templates/extra.qtpl:206
		qw422016.N().S(`)
	if err != nil {
		in.AddError(err)
	}
`)
		//line templates/extra.qtpl:210
	}
	//line templates/extra.qtpl:210
	qw422016.N().S(`
	*t = `)
	//line templates/extra.qtpl:211
	qw422016.N().S(typ)
	//line templates/extra.qtpl:211
	qw422016.N().S(`(`)
	//line templates/extra.qtpl:211
	if parseFunc != "" {
		//line templates/extra.qtpl:211
		qw422016.N().S(`v`)
		//line templates/extra.qtpl:211
	} else {
		//line templates/extra.qtpl:211
		qw422016.N().S(`buf`)
		//line templates/extra.qtpl:211
	}
	//line templates/extra.qtpl:211
	qw422016.N().S(`)
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *`)
	//line templates/extra.qtpl:215
	qw422016.N().S(typ)
	//line templates/extra.qtpl:215
	qw422016.N().S(`) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
`)
//line templates/extra.qtpl:218
}

//line templates/extra.qtpl:218
func WriteExtraFixStringUnmarshaler(qq422016 qtio422016.Writer, typ, parseFunc, extra string) {
	//line templates/extra.qtpl:218
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/extra.qtpl:218
	StreamExtraFixStringUnmarshaler(qw422016, typ, parseFunc, extra)
	//line templates/extra.qtpl:218
	qt422016.ReleaseWriter(qw422016)
//line templates/extra.qtpl:218
}

//line templates/extra.qtpl:218
func ExtraFixStringUnmarshaler(typ, parseFunc, extra string) string {
	//line templates/extra.qtpl:218
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/extra.qtpl:218
	WriteExtraFixStringUnmarshaler(qb422016, typ, parseFunc, extra)
	//line templates/extra.qtpl:218
	qs422016 := string(qb422016.B)
	//line templates/extra.qtpl:218
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/extra.qtpl:218
	return qs422016
//line templates/extra.qtpl:218
}

// ExtraExceptionDetailsTemplate is a special template for the Runtime.ExceptionDetails type that
// defines the standard error interface.

//line templates/extra.qtpl:222
func StreamExtraExceptionDetailsTemplate(qw422016 *qt422016.Writer) {
	//line templates/extra.qtpl:222
	qw422016.N().S(`
// Error satisfies the error interface.
func (e *ExceptionDetails) Error() string {
	// TODO: watch script parsed events and match the ExceptionDetails.ScriptID
	// to the name/location of the actual code and display here
	return fmt.Sprintf("encountered exception '%s' (%d:%d)", e.Text, e.LineNumber, e.ColumnNumber)
}
`)
//line templates/extra.qtpl:229
}

//line templates/extra.qtpl:229
func WriteExtraExceptionDetailsTemplate(qq422016 qtio422016.Writer) {
	//line templates/extra.qtpl:229
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/extra.qtpl:229
	StreamExtraExceptionDetailsTemplate(qw422016)
	//line templates/extra.qtpl:229
	qt422016.ReleaseWriter(qw422016)
//line templates/extra.qtpl:229
}

//line templates/extra.qtpl:229
func ExtraExceptionDetailsTemplate() string {
	//line templates/extra.qtpl:229
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/extra.qtpl:229
	WriteExtraExceptionDetailsTemplate(qb422016)
	//line templates/extra.qtpl:229
	qs422016 := string(qb422016.B)
	//line templates/extra.qtpl:229
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/extra.qtpl:229
	return qs422016
//line templates/extra.qtpl:229
}

// ExtraCDPTypes is the template for additional internal type
// declarations.

//line templates/extra.qtpl:233
func StreamExtraCDPTypes(qw422016 *qt422016.Writer) {
	//line templates/extra.qtpl:233
	qw422016.N().S(`

// Error satisfies the error interface.
func (t ErrorType) Error() string {
	return string(t)
}

// Handler is the common interface for a Chrome Debugging Protocol target.
type Handler interface {
	// SetActive changes the top level frame id.
	SetActive(context.Context, FrameID) error

	// GetRoot returns the root document node for the top level frame.
	GetRoot(context.Context) (*Node, error)

	// WaitFrame waits for a frame to be available.
	WaitFrame(context.Context, FrameID) (*Frame, error)

	// WaitNode waits for a node to be available.
	WaitNode(context.Context, *Frame, NodeID) (*Node, error)

	// Execute executes the specified command using the supplied context and
	// parameters.
	Execute(context.Context, MethodType, easyjson.Marshaler, easyjson.Unmarshaler) error

	// Listen creates a channel that will receive an event for the types
	// specified.
	Listen(...MethodType) <-chan interface{}

	// Release releases a channel returned from Listen.
	Release(<-chan interface{})
}
`)
//line templates/extra.qtpl:265
}

//line templates/extra.qtpl:265
func WriteExtraCDPTypes(qq422016 qtio422016.Writer) {
	//line templates/extra.qtpl:265
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/extra.qtpl:265
	StreamExtraCDPTypes(qw422016)
	//line templates/extra.qtpl:265
	qt422016.ReleaseWriter(qw422016)
//line templates/extra.qtpl:265
}

//line templates/extra.qtpl:265
func ExtraCDPTypes() string {
	//line templates/extra.qtpl:265
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/extra.qtpl:265
	WriteExtraCDPTypes(qb422016)
	//line templates/extra.qtpl:265
	qs422016 := string(qb422016.B)
	//line templates/extra.qtpl:265
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/extra.qtpl:265
	return qs422016
//line templates/extra.qtpl:265
}

// ExtraUtilTemplate generates the decode func for the Message type.

//line templates/extra.qtpl:268
func StreamExtraUtilTemplate(qw422016 *qt422016.Writer, domains []*internal.Domain) {
	//line templates/extra.qtpl:268
	qw422016.N().S(`
type empty struct{}
var emptyVal = &empty{}

// UnmarshalMessage unmarshals the message result or params.
func UnmarshalMessage(msg *cdp.Message) (interface{}, error) {
	var v easyjson.Unmarshaler
	switch msg.Method {`)
	//line templates/extra.qtpl:275
	for _, d := range domains {
		//line templates/extra.qtpl:275
		for _, c := range d.Commands {
			//line templates/extra.qtpl:275
			qw422016.N().S(`
	case cdp.`)
			//line templates/extra.qtpl:276
			qw422016.N().S(c.CommandMethodType(d))
			//line templates/extra.qtpl:276
			qw422016.N().S(`:`)
			//line templates/extra.qtpl:276
			if len(c.Returns) == 0 {
				//line templates/extra.qtpl:276
				qw422016.N().S(`
		return emptyVal, nil`)
				//line templates/extra.qtpl:277
			} else {
				//line templates/extra.qtpl:277
				qw422016.N().S(`
		v = new(`)
				//line templates/extra.qtpl:278
				qw422016.N().S(d.PackageRefName())
				//line templates/extra.qtpl:278
				qw422016.N().S(`.`)
				//line templates/extra.qtpl:278
				qw422016.N().S(c.CommandReturnsType())
				//line templates/extra.qtpl:278
				qw422016.N().S(`)`)
				//line templates/extra.qtpl:278
			}
			//line templates/extra.qtpl:278
			qw422016.N().S(`
	`)
			//line templates/extra.qtpl:279
		}
		//line templates/extra.qtpl:279
		for _, e := range d.Events {
			//line templates/extra.qtpl:279
			qw422016.N().S(`
	case cdp.`)
			//line templates/extra.qtpl:280
			qw422016.N().S(e.EventMethodType(d))
			//line templates/extra.qtpl:280
			qw422016.N().S(`:
		v = new(`)
			//line templates/extra.qtpl:281
			qw422016.N().S(d.PackageRefName())
			//line templates/extra.qtpl:281
			qw422016.N().S(`.`)
			//line templates/extra.qtpl:281
			qw422016.N().S(e.EventType())
			//line templates/extra.qtpl:281
			qw422016.N().S(`)
	`)
			//line templates/extra.qtpl:282
		}
		//line templates/extra.qtpl:282
	}
	//line templates/extra.qtpl:282
	qw422016.N().S(`}

	var buf easyjson.RawMessage
	switch {
	case msg.Params != nil:
		buf = msg.Params

	case msg.Result != nil:
		buf = msg.Result

	default:
		return nil, errors.New("msg missing params or result")
	}

	err := easyjson.Unmarshal(buf, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
`)
//line templates/extra.qtpl:303
}

//line templates/extra.qtpl:303
func WriteExtraUtilTemplate(qq422016 qtio422016.Writer, domains []*internal.Domain) {
	//line templates/extra.qtpl:303
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/extra.qtpl:303
	StreamExtraUtilTemplate(qw422016, domains)
	//line templates/extra.qtpl:303
	qt422016.ReleaseWriter(qw422016)
//line templates/extra.qtpl:303
}

//line templates/extra.qtpl:303
func ExtraUtilTemplate(domains []*internal.Domain) string {
	//line templates/extra.qtpl:303
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/extra.qtpl:303
	WriteExtraUtilTemplate(qb422016, domains)
	//line templates/extra.qtpl:303
	qs422016 := string(qb422016.B)
	//line templates/extra.qtpl:303
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/extra.qtpl:303
	return qs422016
//line templates/extra.qtpl:303
}

//line templates/extra.qtpl:305
func StreamExtraMethodTypeDomainDecoder(qw422016 *qt422016.Writer) {
	//line templates/extra.qtpl:305
	qw422016.N().S(`
// Domain returns the Chrome Debugging Protocol domain of the event or command.
func (t MethodType) Domain() string {
	return string(t[:strings.IndexByte(string(t), '.')])
}
`)
//line templates/extra.qtpl:310
}

//line templates/extra.qtpl:310
func WriteExtraMethodTypeDomainDecoder(qq422016 qtio422016.Writer) {
	//line templates/extra.qtpl:310
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/extra.qtpl:310
	StreamExtraMethodTypeDomainDecoder(qw422016)
	//line templates/extra.qtpl:310
	qt422016.ReleaseWriter(qw422016)
//line templates/extra.qtpl:310
}

//line templates/extra.qtpl:310
func ExtraMethodTypeDomainDecoder() string {
	//line templates/extra.qtpl:310
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/extra.qtpl:310
	WriteExtraMethodTypeDomainDecoder(qb422016)
	//line templates/extra.qtpl:310
	qs422016 := string(qb422016.B)
	//line templates/extra.qtpl:310
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/extra.qtpl:310
	return qs422016
//line templates/extra.qtpl:310
}
