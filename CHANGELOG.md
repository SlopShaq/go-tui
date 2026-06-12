# Changelog

## [0.17.0](https://github.com/SlopShaq/go-tui/compare/v0.16.0...v0.17.0) (2026-06-12)


### Features

* add alternate-scroll mode so the wheel scrolls without mouse capture ([dda7860](https://github.com/SlopShaq/go-tui/commit/dda78600501d18aa9bfaeee88737a4ec5f1e0e17))
* add ApplyDim and FillBlank methods to Buffer ([0accc33](https://github.com/SlopShaq/go-tui/commit/0accc33649d819769261c50ea48c46af5c2fbf47))
* add auto focus highlighting, Tab cycling, key suppression, and bottom-sheet modal example ([d4533e8](https://github.com/SlopShaq/go-tui/commit/d4533e8f014e8f7ac9757fc67c405c13b088eaa1))
* add autoFocus attribute and make focusColor optional for Input and TextArea ([f45f6f5](https://github.com/SlopShaq/go-tui/commit/f45f6f58468ca2ce457c0877c2cc9a2659259eac))
* add borderGradient, focusGradient, and backspace scroll fix for Input and TextArea ([39a855c](https://github.com/SlopShaq/go-tui/commit/39a855cdf349e5acd2524f7e876bb66f94ed9569))
* add Buffer.SetRuneLink to attach hyperlink targets ([d7c415e](https://github.com/SlopShaq/go-tui/commit/d7c415ed194cd23bd6406739e19c95c52a8a4e9f))
* add default Ctrl+Z suspend fallback in key dispatch ([743b7cf](https://github.com/SlopShaq/go-tui/commit/743b7cfbbbcc1de22f3e34c771f924762e7f4690))
* add Events, DispatchEvents, Step methods for manual event loops ([419fcd6](https://github.com/SlopShaq/go-tui/commit/419fcd6b42c9c18e3e724ccfc3d178c80412f055))
* add FlexWrap and AlignContent types with public API and Tailwind classes ([c926df1](https://github.com/SlopShaq/go-tui/commit/c926df10e2794fabc4e651d498fd70291308dba4))
* add FocusRequired flag to KeyPattern and focus-gated binding helpers ([830c400](https://github.com/SlopShaq/go-tui/commit/830c400c2846a6d7ff6971e4ee4958bb920ba2ea))
* add hideScrollbar option for scrollable containers ([0260291](https://github.com/SlopShaq/go-tui/commit/026029172d2e4b131fd92065ee08e7f4cfcb0652))
* add hideScrollbar option for scrollable containers ([15a2ce7](https://github.com/SlopShaq/go-tui/commit/15a2ce75360ad5c17ba46f80b4b36b327b572b95))
* add Hyperlinks capability, default on except TERM=dumb ([b5633e1](https://github.com/SlopShaq/go-tui/commit/b5633e1cd7878e6f3eadbb77499fa110eb11cb56))
* add IsFocused query method to focusManager ([f6c880f](https://github.com/SlopShaq/go-tui/commit/f6c880fe7c7a1645ec7c8ca4b5ae168a970ffef9))
* add KeyCtrlA-Z helpers, fix Ctrl+H parsing, add topic-based debug logging ([b6eea3d](https://github.com/SlopShaq/go-tui/commit/b6eea3da39d70aa97e6968367e892f7a4b0538d1))
* add KeyMatcher interface with On/OnStop/OnFocused constructors ([26c045a](https://github.com/SlopShaq/go-tui/commit/26c045a713e42487041c5dedf0699eda96e54920))
* add Kitty keyboard negotiation to Terminal interface ([3916469](https://github.com/SlopShaq/go-tui/commit/3916469baf93537ffffca543f06ff5a4e896c046))
* add Kitty keyboard protocol escape sequences ([3d85b9d](https://github.com/SlopShaq/go-tui/commit/3d85b9d36a95a9b7151781ac9c5f40896cc03589))
* add Kitty keyboard protocol for Ctrl+H/I/M disambiguation ([91b6877](https://github.com/SlopShaq/go-tui/commit/91b687723f06cc450601ed4775040e45e66e3b7e))
* add KittyKeyboard capability and WithLegacyKeyboard option ([615db9c](https://github.com/SlopShaq/go-tui/commit/615db9cd3e1f4a87707254edf6b08eb2bdf4ad14))
* add Link field to Cell and include it in Equal/Diff ([61003b5](https://github.com/SlopShaq/go-tui/commit/61003b51e50e5f747aa4f08a498c821cf5e4d071))
* add Markdown component with block tree rendering ([97dd1a4](https://github.com/SlopShaq/go-tui/commit/97dd1a45aa8be69f17a7ddb79730630eed7744b6))
* add markdown rendering support (rich text, OSC 8 links, parser, component, gsx) ([f3cc139](https://github.com/SlopShaq/go-tui/commit/f3cc139dd2b11b479f83dec6dcb8af5f9ec2487b))
* add MarkdownTheme and DefaultMarkdownTheme ([1e46923](https://github.com/SlopShaq/go-tui/commit/1e469231a5059fa49afb0a5b1a047bd59308bd9f))
* add Modal component with Render, KeyListener, and MouseListener ([9b7f342](https://github.com/SlopShaq/go-tui/commit/9b7f34249593b8c4f18a2795740bc6d2a441a492))
* add modal support ([f864718](https://github.com/SlopShaq/go-tui/commit/f864718cb3b090bd40547a66868e5b9163d6977f))
* add modal support to code generator ([7830d29](https://github.com/SlopShaq/go-tui/commit/7830d296299e3f3ca73f50588ef1520a12a71406))
* add modal to analyzer known tags and attributes ([05e017d](https://github.com/SlopShaq/go-tui/commit/05e017da7cd8fb0eb39f538d329d0ecbd6f8d703))
* add no-op suspend stubs for Windows ([c6f8412](https://github.com/SlopShaq/go-tui/commit/c6f8412786176a261f11bc572dafd9ba0aca2f5d))
* add OnChange watcher for reactive state effects ([5292530](https://github.com/SlopShaq/go-tui/commit/52925306c7921c03bee66d2dffa70e54dc9f30c9))
* add OnKeyMod helper, fix flex-wrap align-content, update docs ([af20ac4](https://github.com/SlopShaq/go-tui/commit/af20ac42a8a4532b1e176f5e12bfa87398b2606b))
* add onSuspend/onResume fields to App struct ([77a183d](https://github.com/SlopShaq/go-tui/commit/77a183d409e4fa91db7c468729b8fb97c97c348e))
* add Open lifecycle method, make Close idempotent ([1d2362a](https://github.com/SlopShaq/go-tui/commit/1d2362adc189af097bbb83be51ed78321ce63321))
* add OSC 8 hyperlink escape sequences to escBuilder ([0409c6f](https://github.com/SlopShaq/go-tui/commit/0409c6fea3d456fc042970748a16495d1606854a))
* add overlay flag to Element for modal support ([a131933](https://github.com/SlopShaq/go-tui/commit/a1319332c44dd31786b77b00490c1df1243a2633))
* add overlay registration system to App ([af4c29b](https://github.com/SlopShaq/go-tui/commit/af4c29b20d93fe9807f1d8b00ef616dd48869348))
* add overlay render pass and focus scoping for modal support ([cff7051](https://github.com/SlopShaq/go-tui/commit/cff70514dabf63b44a77d8705510a65cfd718de5))
* add pluggable CodeHighlighter with built-in palette ([e1d7080](https://github.com/SlopShaq/go-tui/commit/e1d70800c5a192844234ba85af4385ad3789640b))
* add PostRenderHook and allow invisible virtual cursor ([ac3e9f4](https://github.com/SlopShaq/go-tui/commit/ac3e9f428f5863f7b9341be2daee8f61acbaab4c))
* add PostRenderHook and allow invisible virtual cursor ([bdd0068](https://github.com/SlopShaq/go-tui/commit/bdd006874f378cbcaeb07c15defd271a34399147))
* add Print, Sprint, and Fprint for single-frame rendering ([015ae90](https://github.com/SlopShaq/go-tui/commit/015ae90aa694fefb9801e791f73dca3290857d70))
* add Print, Sprint, and Fprint for single-frame rendering ([b8f34be](https://github.com/SlopShaq/go-tui/commit/b8f34be9df2acc7b4effc7fad051f0efec5c3d93))
* add reactive value binding and focusColor for Input and TextArea ([2e7bb70](https://github.com/SlopShaq/go-tui/commit/2e7bb709d95a0ac0a7c7aee713d4a9aa2540120b))
* add reset confirmation modal to color mixer example ([9308385](https://github.com/SlopShaq/go-tui/commit/930838565ba3ee3646a85991bd64cf4fe6607827))
* add single-frame print example (examples/19-print) ([89b6e60](https://github.com/SlopShaq/go-tui/commit/89b6e608b540da6279fe63b2931ec4a3b8812b60))
* add TextSpan rich-text type with accessors, clearing, and style helpers ([161339b](https://github.com/SlopShaq/go-tui/commit/161339b4ff89533dee042fd0233d4b9e74e98d93))
* add UpdateEvent type for queued closures ([3a85fce](https://github.com/SlopShaq/go-tui/commit/3a85fce643e6fb013e219a94c5051d810c7d173b))
* add WithBorderTitle option for rendering titles in box borders ([da4c10b](https://github.com/SlopShaq/go-tui/commit/da4c10b8953125cb8ca90c6688bea6104edc4aeb))
* add WithBorderTitle option, drawBoxTitle helpers, GSX integration, tests ([31ba6dc](https://github.com/SlopShaq/go-tui/commit/31ba6dc2c27ed69de4b18aeac9aaae1bac25ff6c))
* add WithModalKeyMap and tie key blocking to trapFocus ([#48](https://github.com/SlopShaq/go-tui/issues/48)) ([2dcff81](https://github.com/SlopShaq/go-tui/commit/2dcff816da24182d22b8af97e22d3edf32923508))
* add WithOnSuspend and WithOnResume app options ([cedc55d](https://github.com/SlopShaq/go-tui/commit/cedc55d7b2da7af4ff003973a50a95293d47b7ce))
* add Wrap() public getter to Element ([3bfb4b2](https://github.com/SlopShaq/go-tui/commit/3bfb4b20ff9889f9cb33cd58433c3b21c8706ad4))
* add wrapSpans for style-preserving word wrapping ([dd29b78](https://github.com/SlopShaq/go-tui/commit/dd29b783b57431913d4b40f0469921c8e941fc1d))
* add zero-dependency markdown parser (internal/markdown) ([b7b549f](https://github.com/SlopShaq/go-tui/commit/b7b549f427b81080b5a85bec1ff4131e68d18e05))
* drop @ prefix from control flow keywords ([9b9891c](https://github.com/SlopShaq/go-tui/commit/9b9891c50c7504784955060d9dddea432cf918d8))
* emit bindAppFields helper for method components ([dcf2824](https://github.com/SlopShaq/go-tui/commit/dcf28247a8d3f0c72222dd5868b996849aac4d2b))
* emit OSC 8 hyperlinks for linked cells in Flush ([143c3b8](https://github.com/SlopShaq/go-tui/commit/143c3b8604bb25772bbef3364c660ca141845330))
* emit OSC 8 hyperlinks in bufferRowToANSI ([b7ef736](https://github.com/SlopShaq/go-tui/commit/b7ef7365cbdfa7c204be7d21cdf03b75850b426f))
* **examples:** add 23-event-dump diagnostic for verifying input handling ([c9ecaf5](https://github.com/SlopShaq/go-tui/commit/c9ecaf53fdc6a187f97f8635734176813fb7c997))
* **examples:** add directory tree data model and flatten logic ([337014b](https://github.com/SlopShaq/go-tui/commit/337014b9b37b01473214ebc7ef244fe2aa280541))
* **examples:** add directory tree example ([fb5e20e](https://github.com/SlopShaq/go-tui/commit/fb5e20e906a196e37bb85fdf849d776d90bde183))
* **examples:** add directory tree example skeleton ([aa4f927](https://github.com/SlopShaq/go-tui/commit/aa4f9275281bbb8cf3fed4569c7b04224c620260))
* **examples:** add directory tree keyboard navigation ([d24c585](https://github.com/SlopShaq/go-tui/commit/d24c585881c6cf52bc22a12346febbabe4363386))
* **examples:** add directory tree render and complete example ([0745719](https://github.com/SlopShaq/go-tui/commit/0745719e66a7996c202a0b1a06348d24bce10abc))
* **examples:** add scrolling, overflow clipping, and ancestor path highlighting ([79fb129](https://github.com/SlopShaq/go-tui/commit/79fb12932332b78251f4cb7c458a43ba88bf21b4))
* **examples:** replace hardcoded tree with random generator ([f79a556](https://github.com/SlopShaq/go-tui/commit/f79a556da20870cd81385949f7b26c3e08cef218))
* **examples:** show selected node path at top of directory tree ([0fdfc83](https://github.com/SlopShaq/go-tui/commit/0fdfc8313f00cd0f3133a356dd437b35baa79a0c))
* extend LetBinding AST for component call and expression RHS ([82c4e5d](https://github.com/SlopShaq/go-tui/commit/82c4e5d98ac542b95c02ea362c465495fdfaa5de))
* extract getCoreType helper and refactor type checking functions ([daa2699](https://github.com/SlopShaq/go-tui/commit/daa26993f5f05c5eebcd18dd8dea1822c8059fa3))
* formatter emits bare if/for/else and := instead of @-prefixed syntax ([27a71bc](https://github.com/SlopShaq/go-tui/commit/27a71bcd6bfc98a8a31e26a36c359f5ed55cab17))
* generalize regex patterns and fix whitespace in analyzer ([7aae716](https://github.com/SlopShaq/go-tui/commit/7aae716eb460bd8cd227b20a6b10b21147788500))
* generator handles component call and expression RHS in LetBinding ([162117a](https://github.com/SlopShaq/go-tui/commit/162117a122a5dea8b2727839ade5bb7a7618994e))
* **highlight:** add Bash lexer ([de4382a](https://github.com/SlopShaq/go-tui/commit/de4382a82193a0094f9b7a096a153b15a384daa2))
* **highlight:** add Go lexer ([a906f6e](https://github.com/SlopShaq/go-tui/commit/a906f6ec4e96e509b4143eca4cc20356811c8240))
* **highlight:** add JS/TS lexer ([5a92b50](https://github.com/SlopShaq/go-tui/commit/5a92b50f30812035fbba5a0d246510bb87f476b8))
* **highlight:** add JSON lexer with key/value distinction ([5086351](https://github.com/SlopShaq/go-tui/commit/5086351fcfae314ad94d2dbe293f3ec0dddb18ae))
* **highlight:** add tokenizer core with language dispatch and fallback ([cfe2bc8](https://github.com/SlopShaq/go-tui/commit/cfe2bc881f3b497de0427ab82ad586e598816938))
* implement focus-gated dispatch in dispatch table ([20d8d01](https://github.com/SlopShaq/go-tui/commit/20d8d01345d86c46eedc7eabbd6ba8bb21751874))
* implement Kitty keyboard protocol negotiation in ANSITerminal ([dd3ac92](https://github.com/SlopShaq/go-tui/commit/dd3ac92c8f769ef940d24a584f377a1df5aacc49))
* implement suspend/resume terminal state management ([0fb4125](https://github.com/SlopShaq/go-tui/commit/0fb4125517966c7e52c295b14bab2c11004a0ce6))
* include kitty-keyboard in capabilities string ([87f78e3](https://github.com/SlopShaq/go-tui/commit/87f78e3a961ccf8f45d40a46b52a2f2bc2529ca7))
* **layout:** implement flex-wrap line breaking, per-line layout, and auto cross-axis sizing ([3fed657](https://github.com/SlopShaq/go-tui/commit/3fed657fe979ffb7c47e78b3611112eb8202c2f6))
* lexer and parser accept bare if/for/else, remove TokenAtIf/TokenAtFor/TokenAtElse ([c54e604](https://github.com/SlopShaq/go-tui/commit/c54e604fda74f3cb3d867d02e22c7269fd261f42))
* LSP uses bare keyword syntax in completions, hover, and semantic tokens ([d2107b7](https://github.com/SlopShaq/go-tui/commit/d2107b77a4215a60ba25c45d5898bbdbfa9957c8))
* **lsp:** add modal element definition to schema ([fe5aab6](https://github.com/SlopShaq/go-tui/commit/fe5aab68812fdf6c6cb406f67325427330c0b3ae))
* manual event loop API (Open/Events/Dispatch/Render/Close) ([973bf89](https://github.com/SlopShaq/go-tui/commit/973bf892dcc5780b6d24620f2eb2c19228cc0e3a))
* **markdown:** highlight fenced code blocks via the theme highlighter ([d173a0c](https://github.com/SlopShaq/go-tui/commit/d173a0ce1513cf32a36ef6655cf419caaea78159))
* **markdown:** render tables as a full grid ([0ad97c8](https://github.com/SlopShaq/go-tui/commit/0ad97c896847421b58ab33270102d2006814afd8))
* **markdown:** restyle headings/blockquotes, outline tables ([44e9b60](https://github.com/SlopShaq/go-tui/commit/44e9b606865fc12a13439a61ee6c3361dfdc8d62))
* **markdown:** rule between every table row, blank line after headings ([de8c1f0](https://github.com/SlopShaq/go-tui/commit/de8c1f07ba044cfdd688650c388a187633bc01d9))
* **markdown:** scrollable code blocks, heading spacing, h1/h2 italics ([d50dfbb](https://github.com/SlopShaq/go-tui/commit/d50dfbb706e4aa81928bf9203a952effb5cf9f80))
* **markdown:** wire built-in highlighter into the default theme ([c052b58](https://github.com/SlopShaq/go-tui/commit/c052b583e7fb7e45c70a8fc193d2d375b77f248b))
* measure rich text in IntrinsicSize and HeightForWidth ([63ba8eb](https://github.com/SlopShaq/go-tui/commit/63ba8eb3ea3cb64765ccfc4dbeb3b2f43cf8a7fb))
* parse bare if/for/else and := / var bindings with element/component RHS ([c3c5ac6](https://github.com/SlopShaq/go-tui/commit/c3c5ac6743f7b1e914c9e48ab022c6653474ebf4))
* parse Kitty keyboard protocol CSI u sequences ([d8e902b](https://github.com/SlopShaq/go-tui/commit/d8e902bb704e9ce5d918b3809c31ed0d100e8a11))
* register &lt;markdown&gt; tag in analyzer and LSP schema ([ba5d647](https://github.com/SlopShaq/go-tui/commit/ba5d6476182fe8847a95a4739b9cd975f8190f09))
* register SIGTSTP signal handler in app event loop ([10c42de](https://github.com/SlopShaq/go-tui/commit/10c42de3ac29fcea211918a6cfd7d3586652023d))
* reject @if/@for/@else/[@let](https://github.com/let) with diagnostic errors, remove TokenAtLet and parseLet ([0dae91b](https://github.com/SlopShaq/go-tui/commit/0dae91b0047a341ed122cd5992c0fed1b9adac02))
* remove @-prefixed control flow from tree-sitter grammar ([6bd0d78](https://github.com/SlopShaq/go-tui/commit/6bd0d78c099f6cf5bf784e15ece7175b26cf40c2))
* remove legacy @-prefixed control flow syntax ([cd711cb](https://github.com/SlopShaq/go-tui/commit/cd711cb2e75b325442a59267b283933f6feedc21))
* render code fences, lists, blockquotes, tables; fix row layout ([b58341c](https://github.com/SlopShaq/go-tui/commit/b58341c03cd850246e473272d02761e325c262fd))
* render rich text in the clipped/scroll render pass ([c665816](https://github.com/SlopShaq/go-tui/commit/c66581663cef5f23811db91e0e5b3f2eaad505b0))
* render rich text in the normal pass; fix wrapSpans word boundaries ([2331b45](https://github.com/SlopShaq/go-tui/commit/2331b4588b222a81718e0b9325646ca7222c0b8e))
* separate tabStop from focusable to fix Tab navigation ([f761d8a](https://github.com/SlopShaq/go-tui/commit/f761d8acc17c25ad9133408c4daf3426099cfafd))
* split event channels with input priority fan-in ([60730a0](https://github.com/SlopShaq/go-tui/commit/60730a0d59333cae09bb42293fb1f5ee7c392918))
* support multi-line component call arguments ([c31e6ac](https://github.com/SlopShaq/go-tui/commit/c31e6acfbec24b3009b289b33f67e5881a8c33d6))
* TextMate grammar highlights bare if/for/else and := bindings ([a7e6bb8](https://github.com/SlopShaq/go-tui/commit/a7e6bb8a97baaac17bf2d1acb76193496fce016b))
* thread TextSpan.Link through wrapping and rich-text rendering ([d551098](https://github.com/SlopShaq/go-tui/commit/d55109880fc07ae3caa5d1aa9f69af7a13fb9ee0))
* tree-sitter grammar supports bare if/for/else and := bindings ([f79eee9](https://github.com/SlopShaq/go-tui/commit/f79eee92cf6e8ed241ff49240f7f3f71b6da0d4a))
* trim trailing whitespace so copied selections stay clean ([d89ae47](https://github.com/SlopShaq/go-tui/commit/d89ae47c23bf99b55015cafb63696be5051c7138))
* **tuigen:** generate tui.NewMarkdown for the &lt;markdown&gt; tag ([651d8e0](https://github.com/SlopShaq/go-tui/commit/651d8e06722c9b9ba215e1d41a626858f7cd4dff))
* wire Input component into focus system with focus-gated key bindings ([3a5d28c](https://github.com/SlopShaq/go-tui/commit/3a5d28c11820f38df8e9fce359d220bea6de36ed))
* wire Kitty keyboard negotiation into app lifecycle ([6b836db](https://github.com/SlopShaq/go-tui/commit/6b836db0c22dba34f8b7c6166c790d74ab264dce))
* wire TextArea component into focus system with focus-gated key bindings ([036e91b](https://github.com/SlopShaq/go-tui/commit/036e91bee7d84e57381f5515a14322bcccfb0696))


### Bug Fixes

* add Enter key handler to select focused modal button in color mixer ([324b540](https://github.com/SlopShaq/go-tui/commit/324b54046d9a3c69d6ec933a50042a2a7c257ae7))
* add Tab/Shift+Tab focus navigation to elements example ([ec658d2](https://github.com/SlopShaq/go-tui/commit/ec658d25e67247463533ef43547e08c33ee6e5ea))
* add trailing newline to reader_types.go and document Rune(0) behavior ([b5a15fd](https://github.com/SlopShaq/go-tui/commit/b5a15fd49ddb80a08c47e9c3992a466321aeb340))
* address code review findings ([142f8bb](https://github.com/SlopShaq/go-tui/commit/142f8bb6c87588293e62131624365f125ab65279))
* address code review findings for Kitty keyboard and terminal flush ([4765f41](https://github.com/SlopShaq/go-tui/commit/4765f41d585e7fd2c10c45eada627a2d3d701821))
* address code review findings for Kitty keyboard implementation ([571522f](https://github.com/SlopShaq/go-tui/commit/571522f0bb283a6b813801f8070a887f2a2ab9cb))
* address review feedback for modal PR ([1e7adc4](https://github.com/SlopShaq/go-tui/commit/1e7adc4349411cc3974dc34e623b6113864d0c17))
* address review feedback on debug logging ([110ebed](https://github.com/SlopShaq/go-tui/commit/110ebeda93b711b4dbf8f2417d7520b0645fad32))
* analyzer handles component call and expression RHS in LetBinding ([6c3d46b](https://github.com/SlopShaq/go-tui/commit/6c3d46b9b19d1280d347fb4346869592574604a5))
* apply WithWrap(false) unconditionally ([be350ba](https://github.com/SlopShaq/go-tui/commit/be350ba9037d04feab4160941077467a32dcc8da))
* apply WithWrap(false) unconditionally, fix textarea.go cleanup ([646a21f](https://github.com/SlopShaq/go-tui/commit/646a21fc38872fc245a0ed06782d1ab7051c934e))
* apply WithWrap(false) unconditionally, table-driven map format ([2de5709](https://github.com/SlopShaq/go-tui/commit/2de5709db8e347371b54ef518875f60944ac3151))
* auto-scroll message feed to bottom ([7495525](https://github.com/SlopShaq/go-tui/commit/7495525a494ae2835cb554fc989bd3c51492f512))
* bake inline widget to scrollback before suspend ([75c74e8](https://github.com/SlopShaq/go-tui/commit/75c74e8324e6f02a2e060ebc0edd5bb32a3e7fee))
* **buffer:** keep wide-char continuation cells out of Diff tail erase ([7b77401](https://github.com/SlopShaq/go-tui/commit/7b7740164d93b2b4970298fa9c0695274b869b2d))
* **buffer:** keep wide-char continuation cells out of Diff tail erase ([e5360fb](https://github.com/SlopShaq/go-tui/commit/e5360fbb867512d2729f09faa0aa13a00414a475)), closes [#79](https://github.com/SlopShaq/go-tui/issues/79)
* clear widget area on inline suspend instead of baking duplicate ([650bb33](https://github.com/SlopShaq/go-tui/commit/650bb3382fc712e505c78767d067e89bcb25edf9))
* **docs:** move Google Fonts from CSS [@import](https://github.com/import) to HTML link tag ([d5920ad](https://github.com/SlopShaq/go-tui/commit/d5920ad5b18d4b5adbe444c4f9f01172563f715c))
* **docs:** move playwright to devDependencies and fix tview widget count ([4520a90](https://github.com/SlopShaq/go-tui/commit/4520a9053cc3cc677346f4b2c42989ef56693e29))
* **docs:** remove unused useRef import ([b1c00b9](https://github.com/SlopShaq/go-tui/commit/b1c00b9de1b43dcb356ff00cbd1eba82e27cba3b))
* drain mount cache and unbind outgoing root on root reset ([a3665a9](https://github.com/SlopShaq/go-tui/commit/a3665a9f7ddb97cb309a33771bfd6901dc7c222f))
* drawSpanLines clips an over-wide line instead of dropping the rest ([caeefee](https://github.com/SlopShaq/go-tui/commit/caeefeed638b7e225e055b92c4b82e18fbc42d1e))
* eliminate TOCTOU race in Run/Open and normalize struct alignment ([940ddd9](https://github.com/SlopShaq/go-tui/commit/940ddd90b342a7824efb24969100d73d9616846d))
* end a paragraph when a table starts on the next line ([d1d652d](https://github.com/SlopShaq/go-tui/commit/d1d652d9ac4b1fed100f0717f07d07e0e8e68598))
* ensure modal content children have opaque background by default ([acb4fb1](https://github.com/SlopShaq/go-tui/commit/acb4fb1fd26042114a22e18ddde5d6083ff8dedf))
* **examples:** remove overflow-hidden that bypassed scrollable rendering ([6b38f69](https://github.com/SlopShaq/go-tui/commit/6b38f690ca4b659b654f9f1a705622edb94f249f))
* **examples:** use state-driven scrollOffset for directory tree scrolling ([868021f](https://github.com/SlopShaq/go-tui/commit/868021fb1fe0f2cb67cde9c8c0132a5998325910))
* exclude focus-gated entries from dispatch table conflict validation ([1d264a5](https://github.com/SlopShaq/go-tui/commit/1d264a5071a8e194d3e4a95794546c81a35f7e27))
* flaky tests, nil-ref panic, and fill test coverage gaps ([02f7812](https://github.com/SlopShaq/go-tui/commit/02f781240be99b05a6cda84e7504a0466b322820))
* **formatter:** run gofmt on Go code blocks in .gsx files ([80dcff3](https://github.com/SlopShaq/go-tui/commit/80dcff36605403815056d6e23a09f0a6f744201b))
* gate debug.Log on allTopics so DEBUG=keys doesn't enable untopiced logs ([6ecdfe6](https://github.com/SlopShaq/go-tui/commit/6ecdfe6ec1ae44b2cf5e964aadc2883667b95749))
* gate OSC 8 hyperlinks off for dumb terminals and report them in String() ([781e1cc](https://github.com/SlopShaq/go-tui/commit/781e1ccf674ba3a89eab3291dcc64a762d332d49))
* guard against nil ref in HandleClicks ([8afcd19](https://github.com/SlopShaq/go-tui/commit/8afcd1969101d5228b92e3ea26d26443f2cf0639))
* guard modals in inline mode, fix mouse coords, and resolve dispatch validation ([9313d21](https://github.com/SlopShaq/go-tui/commit/9313d21a9c1fb486030e2a36537dac52559def83))
* handle inline mode suspend/resume without corrupting scrollback ([4e47747](https://github.com/SlopShaq/go-tui/commit/4e477477ae0aba212c90e3b20dce04c8f07a0659))
* handle Kitty keyboard on suspend/resume and clean up minor issues ([8fce79a](https://github.com/SlopShaq/go-tui/commit/8fce79a406b38d8deed91d3957dc9e2156c32e10))
* **highlight:** do not color func receiver or anonymous func params as types ([dbdeb45](https://github.com/SlopShaq/go-tui/commit/dbdeb459379b1c92b31fa0cd004a7afa38f1453e))
* kitty keyboard negotiation and gofmt for .gsx Go code ([08a633f](https://github.com/SlopShaq/go-tui/commit/08a633f495818e41160258dfed1c1f1238967233))
* **layout:** respect explicit WithHeight/WithWidth in IntrinsicSize and HeightForWidth ([7f5f6d6](https://github.com/SlopShaq/go-tui/commit/7f5f6d6e749b68b1ab5aba62821f3051eff398d1))
* **layout:** restrict IntrinsicSize and HeightForWidth overrides to fixed dimensions only ([9327491](https://github.com/SlopShaq/go-tui/commit/93274916870a9e81fae49a088ab4f354580e62ef))
* lineWithCursor checks hideVirtualCursor instead of magic cursorRune value ([8801ced](https://github.com/SlopShaq/go-tui/commit/8801cedb04ab832d6c8a44d597d7e8480a384d09))
* **lsp:** drop @ prefix from offset calculations in LSP providers ([b879fe9](https://github.com/SlopShaq/go-tui/commit/b879fe90bcb70765813ac4a678321c0eb1422bc1))
* **lsp:** highlight Go keywords in expression attribute values ([3e0d00d](https://github.com/SlopShaq/go-tui/commit/3e0d00df48328eb9497f43300e8dc454f001eb3f))
* **lsp:** prevent findElseKeyword false-positive on else inside strings ([f3bbb6f](https://github.com/SlopShaq/go-tui/commit/f3bbb6f0d574e78813ca93862dfbd04e765f9421))
* **lsp:** use character offsets for semantic tokens ([de747fe](https://github.com/SlopShaq/go-tui/commit/de747fe1aade60e15be29f0cb819bc2056db0b11))
* **lsp:** use character offsets instead of byte offsets for semantic tokens ([5106911](https://github.com/SlopShaq/go-tui/commit/5106911db67aff70861a5829926ae99189edbc86))
* **lsp:** wait for gopls to exit before canceling its command context ([464d8b3](https://github.com/SlopShaq/go-tui/commit/464d8b30c68d266fcb17a87bf3bf719d740a6cd5))
* make ContainsPoint account for scroll offset in scrollable containers ([d11f711](https://github.com/SlopShaq/go-tui/commit/d11f711dbcacdfcd2dc5e714790413d3f9f54f95))
* make KeyCtrlH/I/M aliases for KeyBackspace/Tab/Enter ([6907a87](https://github.com/SlopShaq/go-tui/commit/6907a877be042b05363541ee2f26ee554bc76508))
* **markdown:** keep unmatched emphasis delimiters literal ([248109d](https://github.com/SlopShaq/go-tui/commit/248109d7bc62f19cf14d133d732ae6ba098ab687))
* **markdown:** scroll extent, continuous links; responsive example ([e182aeb](https://github.com/SlopShaq/go-tui/commit/e182aebe636a290422a2a2b547d0570aefed5845))
* **markdown:** stop a closing double-marker from italicizing a stray single marker ([09272c0](https://github.com/SlopShaq/go-tui/commit/09272c09962bfb575ee17d4d9a1b29940587baec))
* modal catch-all key binding blocking input when trapFocus is false ([30df011](https://github.com/SlopShaq/go-tui/commit/30df011c10d3dcee507b6779ffdb66f5781f0076))
* **modal:** gate catch-all key binding and focus save/restore on trapFocus ([7dca970](https://github.com/SlopShaq/go-tui/commit/7dca9700bd1b685fda362c0dc6c5900035a4b7fd))
* move SIGWINCH handling from reader to App and default to blocking input ([ef90697](https://github.com/SlopShaq/go-tui/commit/ef9069765ba6ac134ed023b25775c7fe6d616edf))
* negotiate Kitty keyboard after entering alt screen ([7746d43](https://github.com/SlopShaq/go-tui/commit/7746d438eaec2a6dc74c14b7727a3bcb39aaba6d))
* panic on unreachable LetBinding variant, suppress duplicate error for [@let](https://github.com/let) ([4dc7ad3](https://github.com/SlopShaq/go-tui/commit/4dc7ad31e0e0e0158506fda00f62f49708a28bdd))
* parseGoStatement stops at closing brace at depth 0 ([39c5109](https://github.com/SlopShaq/go-tui/commit/39c5109be3497401bfa2a09ab69aeb4180038ee2))
* **parser:** prevent speculative parsing from leaking errors and comments ([9a7d97c](https://github.com/SlopShaq/go-tui/commit/9a7d97ceb06e4b93801770bca749c2c93ac82e35))
* **parser:** use rune-aware iteration for UTF-8 correctness ([cb2a2ae](https://github.com/SlopShaq/go-tui/commit/cb2a2ae2863c2f5177effdcae92e0a6f1936495e))
* preserve consecutive spaces in textarea rendering ([2ebaab6](https://github.com/SlopShaq/go-tui/commit/2ebaab677d86e246d8e5b5efdca3ce34ce224b91))
* preserve consecutive spaces in wrapParagraph ([80a9faf](https://github.com/SlopShaq/go-tui/commit/80a9faf34ae06448a2b6ea07b318717babd242c4))
* preserve EraseToEOL in inline renderer coordinate translation ([5b79a4f](https://github.com/SlopShaq/go-tui/commit/5b79a4f4f32d4f9bdce68019c8803027a933fc81))
* preserve EraseToEOL in inline renderer coordinate translation ([49e2fd0](https://github.com/SlopShaq/go-tui/commit/49e2fd03233f6042475cf18300577e5286b08fb0))
* prevent stack overflow from circular state dependencies ([71f1780](https://github.com/SlopShaq/go-tui/commit/71f1780c08972463805d536cb224614dfd700986))
* re-register SIGTSTP signal handler after resume ([4003004](https://github.com/SlopShaq/go-tui/commit/4003004fe22e803e1bc62fba271a7582c0a6342d))
* refine example text ([3826b5a](https://github.com/SlopShaq/go-tui/commit/3826b5a043c0ebbc10fa6b93c617984fc93c337a))
* remove ability to accidentally generate incorrect Watcher signature ([d34e5b0](https://github.com/SlopShaq/go-tui/commit/d34e5b0c545054f5c6c0e08659be4bdcccebaa45))
* remove flag-2+ dead code from kittySpecialKeys and harden Kitty negotiation ([e897dd1](https://github.com/SlopShaq/go-tui/commit/e897dd19a99e4505ea3b0a1c5bc3574d7368692a))
* repair RenderFull overlay handling and modal inline mode key blocking ([2fc7927](https://github.com/SlopShaq/go-tui/commit/2fc79274fbf8c4d727fad4173093e93b18ba19db))
* replace HTML-style comments with Go-style comments in docs code blocks ([48ed11b](https://github.com/SlopShaq/go-tui/commit/48ed11b5f5c1b485ced35d4358d6e4e08e4eb093))
* replace sleep-based test sync with done channel, document FlexGrow heuristic ([a1db92f](https://github.com/SlopShaq/go-tui/commit/a1db92fa1a98c769f0176bb68b3fcbc80882a1dc))
* replace time.Sleep with channel draining in watcher tests ([3d85b0b](https://github.com/SlopShaq/go-tui/commit/3d85b0b86d6bd747b075c9b73488c3f3000f47b4))
* report initLocked errors to stderr instead of silently dropping logs ([e3f5383](https://github.com/SlopShaq/go-tui/commit/e3f53835851ea130f35f300dc820ce15bad5e45d))
* resolve render consistency, focus lifecycle, and test robustness ([2547181](https://github.com/SlopShaq/go-tui/commit/2547181b1c984fb4a9e8f98a1a86fcc203fff352))
* resolve three suspend/resume issues ([f826211](https://github.com/SlopShaq/go-tui/commit/f82621104df799fc25c858877e60efb0fcdc1958))
* restore deleted tests, scope WithWrap to hideVirtualCursor, table-driven format ([a910d60](https://github.com/SlopShaq/go-tui/commit/a910d60130146a87d43216e213691f7c99c9d52b))
* scope WithWrap to hideVirtualCursor, fix test guard ([e1b8e38](https://github.com/SlopShaq/go-tui/commit/e1b8e387440dc9ead00a3dee316fb5a3f378b248))
* silent event delivery loss after root reset ([378e107](https://github.com/SlopShaq/go-tui/commit/378e10797e7e799a9e5c207fbe33425ea8e8b22e))
* silent interface failures and layout dimension bugs ([8de39e7](https://github.com/SlopShaq/go-tui/commit/8de39e7af675f2bdd5b48d57332e1559803254f6))
* stop background goroutines on NewApp error path ([10d630d](https://github.com/SlopShaq/go-tui/commit/10d630db63a751df5a58c115375e7fea52b7c237))
* struct alignment via gofmt, smarter response terminator detection ([c04060a](https://github.com/SlopShaq/go-tui/commit/c04060ac97b6f5e2d917d1c3c13f263b42c289ff))
* **suspend:** prevent Kitty keyboard response from leaking as typed input after Ctrl+Z resume ([54b4716](https://github.com/SlopShaq/go-tui/commit/54b471612ab13f6f3c5e120965fd9f6a01739e76))
* **test:** append .exe to integration test binary path on Windows ([0af9408](https://github.com/SlopShaq/go-tui/commit/0af940859314b46001cb8c5ce5acbe1fecefde42))
* **textarea:** guard moveEnd at phantom row and overlay cursor on hard-newline boundary ([ddf4f55](https://github.com/SlopShaq/go-tui/commit/ddf4f55658d6bae4987d6cd5da29d519f3217bee))
* **textarea:** keep cursor visible at display-full wrap boundaries ([c2da185](https://github.com/SlopShaq/go-tui/commit/c2da185c4753c221b5ff15122a5727d1875d501c))
* **textarea:** keep the cursor visible at display-full wrap boundaries ([3a2c768](https://github.com/SlopShaq/go-tui/commit/3a2c76812ac1c551e45688a241b2c5cf6e4d485a))
* **textarea:** wrap wide characters at display-width boundaries ([61fb8c8](https://github.com/SlopShaq/go-tui/commit/61fb8c8a5dc548b45574cea6135c9ae95cd9c3e0))
* **textarea:** wrap wide characters at the display-width boundary ([c385312](https://github.com/SlopShaq/go-tui/commit/c385312853f4b5c9a5ddd97d3c753b0bb5a8ba66))
* track outgoing root in a dedicated field for unbind ([e28242b](https://github.com/SlopShaq/go-tui/commit/e28242bdbe0cdcce65fb7fceb46976ee881e9987))
* treat CR/VT/FF as word separators in wrapSpans ([b229624](https://github.com/SlopShaq/go-tui/commit/b229624c2289e12cc259e1a2759cf1b8419e143f))
* **tuigen:** access .Root when method templ root is a function-component call ([976d3c1](https://github.com/SlopShaq/go-tui/commit/976d3c13184856083a250fc872ab995462da3bbc))
* **tuigen:** access .Root when method templ root is a function-component call ([7a61cda](https://github.com/SlopShaq/go-tui/commit/7a61cda34977dcf76f0d3bfe6cfb9a1bdff4e071))
* **tuigen:** adjust source map entries after splicing view reset lines ([c3f2de7](https://github.com/SlopShaq/go-tui/commit/c3f2de7b6d385457033112e8141b38b7c6f79970))
* **tuigen:** avoid receiver shadowing in generated UpdateProps ([6912b40](https://github.com/SlopShaq/go-tui/commit/6912b4086e7f38497f60bf6a26b2c49832e0964c))
* **tuigen:** avoid receiver shadowing in generated UpdateProps ([57378db](https://github.com/SlopShaq/go-tui/commit/57378db6a418b2d4bdcbd0b802c05f729709ca4c))
* **tuigen:** collect all for-loop iteration component views for watcher/bind aggregation ([9309940](https://github.com/SlopShaq/go-tui/commit/9309940e79e45a4828bf487e67e0f872cd2db71d))
* **tuigen:** collect all for-loop iteration component views for watcher/bind aggregation ([2525476](https://github.com/SlopShaq/go-tui/commit/252547614d78fb4760147be42db0564613f723b2))
* **tuigen:** decouple UnbindApp suppression from BindApp detection ([4bdb40d](https://github.com/SlopShaq/go-tui/commit/4bdb40deedec56c1a47553aa6ba4771a4eb61f31))
* **tuigen:** document intentionally ignored inForLoop parameter ([89d7f92](https://github.com/SlopShaq/go-tui/commit/89d7f922170d647ce88dea23e21ba76c53008a50))
* **tuigen:** handle let bindings with call or expression RHS and keep component expression comments ([f0bb13a](https://github.com/SlopShaq/go-tui/commit/f0bb13aa97e5cd1f2bc39dd6bf6830d495e13ade))
* **tuigen:** hoist block-scoped component vars to function scope ([0949b63](https://github.com/SlopShaq/go-tui/commit/0949b63e531c873f4a243885406d032c1b3a7321))
* **tuigen:** hoist conditional and loop-scoped component vars to function scope ([234748b](https://github.com/SlopShaq/go-tui/commit/234748be7925307bdc009fc0e83a5b757ab983bb)), closes [#42](https://github.com/SlopShaq/go-tui/issues/42)
* **tuigen:** only unwrap .Root for function-templ root calls ([cb713b0](https://github.com/SlopShaq/go-tui/commit/cb713b0006dba9d809b4fb946c6ebdaa1dc63627))
* **tuigen:** parenthesize nested loop mount key expressions so multipliers compound ([53765bb](https://github.com/SlopShaq/go-tui/commit/53765bbef6f350e20f9d4def3c3273db8cc8a86b))
* **tuigen:** prevent loop mount keys from colliding with other mount call sites ([d823710](https://github.com/SlopShaq/go-tui/commit/d823710174755e5238a28b4250e07662bf1e90e0))
* **tuigen:** prevent loop mount keys from colliding with standalone component keys ([cda2c0a](https://github.com/SlopShaq/go-tui/commit/cda2c0a02290644e1da922afb712a3979919f7ae))
* **tuigen:** reset views slice in reactive closures to prevent unbounded growth ([89749ce](https://github.com/SlopShaq/go-tui/commit/89749ce4a518ae0fb198a2cca26c4b27767acb36))
* **tuigen:** skip hoisting for method templs with conditional component calls ([b305684](https://github.com/SlopShaq/go-tui/commit/b305684f3928196d7ee3795c94dd4294e0898315))
* update image URLs in example docs ([79152f7](https://github.com/SlopShaq/go-tui/commit/79152f7a80ec54515474a3befbc0adaa6a59b1fa))
* update image URLs in example docs ([7dc88fd](https://github.com/SlopShaq/go-tui/commit/7dc88fd49bc13b6a720b4cbbd9cfa961538441f6))
* update README badges, examples, and VS Code extension docs ([09c8bf6](https://github.com/SlopShaq/go-tui/commit/09c8bf686966539b0b7f17843ece00a7063f9366))
* update stale comments from cursorRune check to hideVirtualCursor ([15ccbe2](https://github.com/SlopShaq/go-tui/commit/15ccbe267d6b82fdaf379932c0c6b4d443e4b601))
* use proper sticky-scroll pattern for message feed ([bf79eb7](https://github.com/SlopShaq/go-tui/commit/bf79eb7987c185180c09f87e90746b56121d7e88))
* use runtime/debug to report version from go install ([438b3f5](https://github.com/SlopShaq/go-tui/commit/438b3f546ef59204da72bbfc3a9ba042fed07c8e))
* **windows:** apply code review feedback on reader ([2529945](https://github.com/SlopShaq/go-tui/commit/2529945dcb2722cf5d5fc5701d6f7e54390497e4))
* **windows:** decode console input records directly for keys, mouse, and resize ([389d343](https://github.com/SlopShaq/go-tui/commit/389d343e0a5dab7122db403e4a734a69d574b19b))
* **windows:** rewrite console reader for keys, mouse, and resize ([da4f753](https://github.com/SlopShaq/go-tui/commit/da4f7539afd6ee6647352b2ec0b0302f74699a09))
* wrap blockquote and list content instead of clipping ([fd9854c](https://github.com/SlopShaq/go-tui/commit/fd9854c9f5ccfd25f949363f86d7f706a242ceb6))


### Performance Improvements

* **highlight:** avoid per-identifier rune slice allocation ([88ff313](https://github.com/SlopShaq/go-tui/commit/88ff313fad66e6f426e4357653465f9c8b828277))

## [0.16.0](https://github.com/grindlemire/go-tui/compare/v0.15.1...v0.16.0) (2026-06-07)


### Features

* add WithBorderTitle option for rendering titles in box borders ([da4c10b](https://github.com/grindlemire/go-tui/commit/da4c10b8953125cb8ca90c6688bea6104edc4aeb))
* add WithBorderTitle option, drawBoxTitle helpers, GSX integration, tests ([31ba6dc](https://github.com/grindlemire/go-tui/commit/31ba6dc2c27ed69de4b18aeac9aaae1bac25ff6c))


### Bug Fixes

* **buffer:** keep wide-char continuation cells out of Diff tail erase ([7b77401](https://github.com/grindlemire/go-tui/commit/7b7740164d93b2b4970298fa9c0695274b869b2d))
* **buffer:** keep wide-char continuation cells out of Diff tail erase ([e5360fb](https://github.com/grindlemire/go-tui/commit/e5360fbb867512d2729f09faa0aa13a00414a475)), closes [#79](https://github.com/grindlemire/go-tui/issues/79)

## [0.15.1](https://github.com/grindlemire/go-tui/compare/v0.15.0...v0.15.1) (2026-06-05)


### Bug Fixes

* preserve EraseToEOL in inline renderer coordinate translation ([5b79a4f](https://github.com/grindlemire/go-tui/commit/5b79a4f4f32d4f9bdce68019c8803027a933fc81))
* preserve EraseToEOL in inline renderer coordinate translation ([49e2fd0](https://github.com/grindlemire/go-tui/commit/49e2fd03233f6042475cf18300577e5286b08fb0))

## [0.15.0](https://github.com/grindlemire/go-tui/compare/v0.14.0...v0.15.0) (2026-06-03)


### Features

* add alternate-scroll mode so the wheel scrolls without mouse capture ([dda7860](https://github.com/grindlemire/go-tui/commit/dda78600501d18aa9bfaeee88737a4ec5f1e0e17))
* add Buffer.SetRuneLink to attach hyperlink targets ([d7c415e](https://github.com/grindlemire/go-tui/commit/d7c415ed194cd23bd6406739e19c95c52a8a4e9f))
* add Hyperlinks capability, default on except TERM=dumb ([b5633e1](https://github.com/grindlemire/go-tui/commit/b5633e1cd7878e6f3eadbb77499fa110eb11cb56))
* add Link field to Cell and include it in Equal/Diff ([61003b5](https://github.com/grindlemire/go-tui/commit/61003b51e50e5f747aa4f08a498c821cf5e4d071))
* add Markdown component with block tree rendering ([97dd1a4](https://github.com/grindlemire/go-tui/commit/97dd1a45aa8be69f17a7ddb79730630eed7744b6))
* add markdown rendering support (rich text, OSC 8 links, parser, component, gsx) ([f3cc139](https://github.com/grindlemire/go-tui/commit/f3cc139dd2b11b479f83dec6dcb8af5f9ec2487b))
* add MarkdownTheme and DefaultMarkdownTheme ([1e46923](https://github.com/grindlemire/go-tui/commit/1e469231a5059fa49afb0a5b1a047bd59308bd9f))
* add OSC 8 hyperlink escape sequences to escBuilder ([0409c6f](https://github.com/grindlemire/go-tui/commit/0409c6fea3d456fc042970748a16495d1606854a))
* add pluggable CodeHighlighter with built-in palette ([e1d7080](https://github.com/grindlemire/go-tui/commit/e1d70800c5a192844234ba85af4385ad3789640b))
* add PostRenderHook and allow invisible virtual cursor ([ac3e9f4](https://github.com/grindlemire/go-tui/commit/ac3e9f428f5863f7b9341be2daee8f61acbaab4c))
* add PostRenderHook and allow invisible virtual cursor ([bdd0068](https://github.com/grindlemire/go-tui/commit/bdd006874f378cbcaeb07c15defd271a34399147))
* add TextSpan rich-text type with accessors, clearing, and style helpers ([161339b](https://github.com/grindlemire/go-tui/commit/161339b4ff89533dee042fd0233d4b9e74e98d93))
* add wrapSpans for style-preserving word wrapping ([dd29b78](https://github.com/grindlemire/go-tui/commit/dd29b783b57431913d4b40f0469921c8e941fc1d))
* add zero-dependency markdown parser (internal/markdown) ([b7b549f](https://github.com/grindlemire/go-tui/commit/b7b549f427b81080b5a85bec1ff4131e68d18e05))
* emit OSC 8 hyperlinks for linked cells in Flush ([143c3b8](https://github.com/grindlemire/go-tui/commit/143c3b8604bb25772bbef3364c660ca141845330))
* emit OSC 8 hyperlinks in bufferRowToANSI ([b7ef736](https://github.com/grindlemire/go-tui/commit/b7ef7365cbdfa7c204be7d21cdf03b75850b426f))
* extract getCoreType helper and refactor type checking functions ([daa2699](https://github.com/grindlemire/go-tui/commit/daa26993f5f05c5eebcd18dd8dea1822c8059fa3))
* generalize regex patterns and fix whitespace in analyzer ([7aae716](https://github.com/grindlemire/go-tui/commit/7aae716eb460bd8cd227b20a6b10b21147788500))
* **highlight:** add Bash lexer ([de4382a](https://github.com/grindlemire/go-tui/commit/de4382a82193a0094f9b7a096a153b15a384daa2))
* **highlight:** add Go lexer ([a906f6e](https://github.com/grindlemire/go-tui/commit/a906f6ec4e96e509b4143eca4cc20356811c8240))
* **highlight:** add JS/TS lexer ([5a92b50](https://github.com/grindlemire/go-tui/commit/5a92b50f30812035fbba5a0d246510bb87f476b8))
* **highlight:** add JSON lexer with key/value distinction ([5086351](https://github.com/grindlemire/go-tui/commit/5086351fcfae314ad94d2dbe293f3ec0dddb18ae))
* **highlight:** add tokenizer core with language dispatch and fallback ([cfe2bc8](https://github.com/grindlemire/go-tui/commit/cfe2bc881f3b497de0427ab82ad586e598816938))
* **markdown:** highlight fenced code blocks via the theme highlighter ([d173a0c](https://github.com/grindlemire/go-tui/commit/d173a0ce1513cf32a36ef6655cf419caaea78159))
* **markdown:** render tables as a full grid ([0ad97c8](https://github.com/grindlemire/go-tui/commit/0ad97c896847421b58ab33270102d2006814afd8))
* **markdown:** restyle headings/blockquotes, outline tables ([44e9b60](https://github.com/grindlemire/go-tui/commit/44e9b606865fc12a13439a61ee6c3361dfdc8d62))
* **markdown:** rule between every table row, blank line after headings ([de8c1f0](https://github.com/grindlemire/go-tui/commit/de8c1f07ba044cfdd688650c388a187633bc01d9))
* **markdown:** scrollable code blocks, heading spacing, h1/h2 italics ([d50dfbb](https://github.com/grindlemire/go-tui/commit/d50dfbb706e4aa81928bf9203a952effb5cf9f80))
* **markdown:** wire built-in highlighter into the default theme ([c052b58](https://github.com/grindlemire/go-tui/commit/c052b583e7fb7e45c70a8fc193d2d375b77f248b))
* measure rich text in IntrinsicSize and HeightForWidth ([63ba8eb](https://github.com/grindlemire/go-tui/commit/63ba8eb3ea3cb64765ccfc4dbeb3b2f43cf8a7fb))
* register &lt;markdown&gt; tag in analyzer and LSP schema ([ba5d647](https://github.com/grindlemire/go-tui/commit/ba5d6476182fe8847a95a4739b9cd975f8190f09))
* render code fences, lists, blockquotes, tables; fix row layout ([b58341c](https://github.com/grindlemire/go-tui/commit/b58341c03cd850246e473272d02761e325c262fd))
* render rich text in the clipped/scroll render pass ([c665816](https://github.com/grindlemire/go-tui/commit/c66581663cef5f23811db91e0e5b3f2eaad505b0))
* render rich text in the normal pass; fix wrapSpans word boundaries ([2331b45](https://github.com/grindlemire/go-tui/commit/2331b4588b222a81718e0b9325646ca7222c0b8e))
* thread TextSpan.Link through wrapping and rich-text rendering ([d551098](https://github.com/grindlemire/go-tui/commit/d55109880fc07ae3caa5d1aa9f69af7a13fb9ee0))
* trim trailing whitespace so copied selections stay clean ([d89ae47](https://github.com/grindlemire/go-tui/commit/d89ae47c23bf99b55015cafb63696be5051c7138))
* **tuigen:** generate tui.NewMarkdown for the &lt;markdown&gt; tag ([651d8e0](https://github.com/grindlemire/go-tui/commit/651d8e06722c9b9ba215e1d41a626858f7cd4dff))


### Bug Fixes

* apply WithWrap(false) unconditionally ([be350ba](https://github.com/grindlemire/go-tui/commit/be350ba9037d04feab4160941077467a32dcc8da))
* apply WithWrap(false) unconditionally, fix textarea.go cleanup ([646a21f](https://github.com/grindlemire/go-tui/commit/646a21fc38872fc245a0ed06782d1ab7051c934e))
* apply WithWrap(false) unconditionally, table-driven map format ([2de5709](https://github.com/grindlemire/go-tui/commit/2de5709db8e347371b54ef518875f60944ac3151))
* drawSpanLines clips an over-wide line instead of dropping the rest ([caeefee](https://github.com/grindlemire/go-tui/commit/caeefeed638b7e225e055b92c4b82e18fbc42d1e))
* end a paragraph when a table starts on the next line ([d1d652d](https://github.com/grindlemire/go-tui/commit/d1d652d9ac4b1fed100f0717f07d07e0e8e68598))
* gate OSC 8 hyperlinks off for dumb terminals and report them in String() ([781e1cc](https://github.com/grindlemire/go-tui/commit/781e1ccf674ba3a89eab3291dcc64a762d332d49))
* **highlight:** do not color func receiver or anonymous func params as types ([dbdeb45](https://github.com/grindlemire/go-tui/commit/dbdeb459379b1c92b31fa0cd004a7afa38f1453e))
* lineWithCursor checks hideVirtualCursor instead of magic cursorRune value ([8801ced](https://github.com/grindlemire/go-tui/commit/8801cedb04ab832d6c8a44d597d7e8480a384d09))
* **markdown:** keep unmatched emphasis delimiters literal ([248109d](https://github.com/grindlemire/go-tui/commit/248109d7bc62f19cf14d133d732ae6ba098ab687))
* **markdown:** scroll extent, continuous links; responsive example ([e182aeb](https://github.com/grindlemire/go-tui/commit/e182aebe636a290422a2a2b547d0570aefed5845))
* **markdown:** stop a closing double-marker from italicizing a stray single marker ([09272c0](https://github.com/grindlemire/go-tui/commit/09272c09962bfb575ee17d4d9a1b29940587baec))
* preserve consecutive spaces in textarea rendering ([2ebaab6](https://github.com/grindlemire/go-tui/commit/2ebaab677d86e246d8e5b5efdca3ce34ce224b91))
* preserve consecutive spaces in wrapParagraph ([80a9faf](https://github.com/grindlemire/go-tui/commit/80a9faf34ae06448a2b6ea07b318717babd242c4))
* restore deleted tests, scope WithWrap to hideVirtualCursor, table-driven format ([a910d60](https://github.com/grindlemire/go-tui/commit/a910d60130146a87d43216e213691f7c99c9d52b))
* scope WithWrap to hideVirtualCursor, fix test guard ([e1b8e38](https://github.com/grindlemire/go-tui/commit/e1b8e387440dc9ead00a3dee316fb5a3f378b248))
* treat CR/VT/FF as word separators in wrapSpans ([b229624](https://github.com/grindlemire/go-tui/commit/b229624c2289e12cc259e1a2759cf1b8419e143f))
* update stale comments from cursorRune check to hideVirtualCursor ([15ccbe2](https://github.com/grindlemire/go-tui/commit/15ccbe267d6b82fdaf379932c0c6b4d443e4b601))
* wrap blockquote and list content instead of clipping ([fd9854c](https://github.com/grindlemire/go-tui/commit/fd9854c9f5ccfd25f949363f86d7f706a242ceb6))


### Performance Improvements

* **highlight:** avoid per-identifier rune slice allocation ([88ff313](https://github.com/grindlemire/go-tui/commit/88ff313fad66e6f426e4357653465f9c8b828277))

## [0.14.0](https://github.com/grindlemire/go-tui/compare/v0.13.1...v0.14.0) (2026-05-21)


### Features

* **examples:** add 23-event-dump diagnostic for verifying input handling ([c9ecaf5](https://github.com/grindlemire/go-tui/commit/c9ecaf53fdc6a187f97f8635734176813fb7c997))


### Bug Fixes

* refine example text ([3826b5a](https://github.com/grindlemire/go-tui/commit/3826b5a043c0ebbc10fa6b93c617984fc93c337a))
* **test:** append .exe to integration test binary path on Windows ([0af9408](https://github.com/grindlemire/go-tui/commit/0af940859314b46001cb8c5ce5acbe1fecefde42))
* **windows:** apply code review feedback on reader ([2529945](https://github.com/grindlemire/go-tui/commit/2529945dcb2722cf5d5fc5701d6f7e54390497e4))
* **windows:** decode console input records directly for keys, mouse, and resize ([389d343](https://github.com/grindlemire/go-tui/commit/389d343e0a5dab7122db403e4a734a69d574b19b))
* **windows:** rewrite console reader for keys, mouse, and resize ([da4f753](https://github.com/grindlemire/go-tui/commit/da4f7539afd6ee6647352b2ec0b0302f74699a09))

## [0.13.1](https://github.com/grindlemire/go-tui/compare/v0.13.0...v0.13.1) (2026-04-24)


### Bug Fixes

* **tuigen:** access .Root when method templ root is a function-component call ([976d3c1](https://github.com/grindlemire/go-tui/commit/976d3c13184856083a250fc872ab995462da3bbc))
* **tuigen:** access .Root when method templ root is a function-component call ([7a61cda](https://github.com/grindlemire/go-tui/commit/7a61cda34977dcf76f0d3bfe6cfb9a1bdff4e071))
* **tuigen:** only unwrap .Root for function-templ root calls ([cb713b0](https://github.com/grindlemire/go-tui/commit/cb713b0006dba9d809b4fb946c6ebdaa1dc63627))

## [0.13.0](https://github.com/grindlemire/go-tui/compare/v0.12.0...v0.13.0) (2026-04-18)


### Features

* add hideScrollbar option for scrollable containers ([15a2ce7](https://github.com/grindlemire/go-tui/commit/15a2ce75360ad5c17ba46f80b4b36b327b572b95))


### Bug Fixes

* update image URLs in example docs ([7dc88fd](https://github.com/grindlemire/go-tui/commit/7dc88fd49bc13b6a720b4cbbd9cfa961538441f6))

## [0.12.0](https://github.com/grindlemire/go-tui/compare/v0.11.0...v0.12.0) (2026-04-18)


### Features

* emit bindAppFields helper for method components ([dcf2824](https://github.com/grindlemire/go-tui/commit/dcf28247a8d3f0c72222dd5868b996849aac4d2b))


### Bug Fixes

* drain mount cache and unbind outgoing root on root reset ([a3665a9](https://github.com/grindlemire/go-tui/commit/a3665a9f7ddb97cb309a33771bfd6901dc7c222f))
* silent event delivery loss after root reset ([378e107](https://github.com/grindlemire/go-tui/commit/378e10797e7e799a9e5c207fbe33425ea8e8b22e))
* track outgoing root in a dedicated field for unbind ([e28242b](https://github.com/grindlemire/go-tui/commit/e28242bdbe0cdcce65fb7fceb46976ee881e9987))
* **tuigen:** decouple UnbindApp suppression from BindApp detection ([4bdb40d](https://github.com/grindlemire/go-tui/commit/4bdb40deedec56c1a47553aa6ba4771a4eb61f31))

## [0.11.0](https://github.com/grindlemire/go-tui/compare/v0.10.3...v0.11.0) (2026-04-03)


### Features

* add WithModalKeyMap and tie key blocking to trapFocus ([#48](https://github.com/grindlemire/go-tui/issues/48)) ([2dcff81](https://github.com/grindlemire/go-tui/commit/2dcff816da24182d22b8af97e22d3edf32923508))

## [0.10.3](https://github.com/grindlemire/go-tui/compare/v0.10.2...v0.10.3) (2026-04-03)


### Bug Fixes

* modal catch-all key binding blocking input when trapFocus is false ([30df011](https://github.com/grindlemire/go-tui/commit/30df011c10d3dcee507b6779ffdb66f5781f0076))
* **modal:** gate catch-all key binding and focus save/restore on trapFocus ([7dca970](https://github.com/grindlemire/go-tui/commit/7dca9700bd1b685fda362c0dc6c5900035a4b7fd))

## [0.10.2](https://github.com/grindlemire/go-tui/compare/v0.10.1...v0.10.2) (2026-04-02)


### Bug Fixes

* **layout:** respect explicit WithHeight/WithWidth in IntrinsicSize and HeightForWidth ([7f5f6d6](https://github.com/grindlemire/go-tui/commit/7f5f6d6e749b68b1ab5aba62821f3051eff398d1))
* **layout:** restrict IntrinsicSize and HeightForWidth overrides to fixed dimensions only ([9327491](https://github.com/grindlemire/go-tui/commit/93274916870a9e81fae49a088ab4f354580e62ef))
* remove ability to accidentally generate incorrect Watcher signature ([d34e5b0](https://github.com/grindlemire/go-tui/commit/d34e5b0c545054f5c6c0e08659be4bdcccebaa45))
* silent interface failures and layout dimension bugs ([8de39e7](https://github.com/grindlemire/go-tui/commit/8de39e7af675f2bdd5b48d57332e1559803254f6))
* **tuigen:** adjust source map entries after splicing view reset lines ([c3f2de7](https://github.com/grindlemire/go-tui/commit/c3f2de7b6d385457033112e8141b38b7c6f79970))
* **tuigen:** collect all for-loop iteration component views for watcher/bind aggregation ([9309940](https://github.com/grindlemire/go-tui/commit/9309940e79e45a4828bf487e67e0f872cd2db71d))
* **tuigen:** collect all for-loop iteration component views for watcher/bind aggregation ([2525476](https://github.com/grindlemire/go-tui/commit/252547614d78fb4760147be42db0564613f723b2))
* **tuigen:** document intentionally ignored inForLoop parameter ([89d7f92](https://github.com/grindlemire/go-tui/commit/89d7f922170d647ce88dea23e21ba76c53008a50))
* **tuigen:** hoist block-scoped component vars to function scope ([0949b63](https://github.com/grindlemire/go-tui/commit/0949b63e531c873f4a243885406d032c1b3a7321))
* **tuigen:** hoist conditional and loop-scoped component vars to function scope ([234748b](https://github.com/grindlemire/go-tui/commit/234748be7925307bdc009fc0e83a5b757ab983bb)), closes [#42](https://github.com/grindlemire/go-tui/issues/42)
* **tuigen:** reset views slice in reactive closures to prevent unbounded growth ([89749ce](https://github.com/grindlemire/go-tui/commit/89749ce4a518ae0fb198a2cca26c4b27767acb36))
* **tuigen:** skip hoisting for method templs with conditional component calls ([b305684](https://github.com/grindlemire/go-tui/commit/b305684f3928196d7ee3795c94dd4294e0898315))

## [0.10.1](https://github.com/grindlemire/go-tui/compare/v0.10.0...v0.10.1) (2026-03-28)


### Bug Fixes

* **formatter:** run gofmt on Go code blocks in .gsx files ([80dcff3](https://github.com/grindlemire/go-tui/commit/80dcff36605403815056d6e23a09f0a6f744201b))
* kitty keyboard negotiation and gofmt for .gsx Go code ([08a633f](https://github.com/grindlemire/go-tui/commit/08a633f495818e41160258dfed1c1f1238967233))
* negotiate Kitty keyboard after entering alt screen ([7746d43](https://github.com/grindlemire/go-tui/commit/7746d438eaec2a6dc74c14b7727a3bcb39aaba6d))

## [0.10.0](https://github.com/grindlemire/go-tui/compare/v0.9.0...v0.10.0) (2026-03-26)


### Features

* add KeyCtrlA-Z helpers, fix Ctrl+H parsing, add topic-based debug logging ([b6eea3d](https://github.com/grindlemire/go-tui/commit/b6eea3da39d70aa97e6968367e892f7a4b0538d1))


### Bug Fixes

* address review feedback on debug logging ([110ebed](https://github.com/grindlemire/go-tui/commit/110ebeda93b711b4dbf8f2417d7520b0645fad32))
* gate debug.Log on allTopics so DEBUG=keys doesn't enable untopiced logs ([6ecdfe6](https://github.com/grindlemire/go-tui/commit/6ecdfe6ec1ae44b2cf5e964aadc2883667b95749))
* report initLocked errors to stderr instead of silently dropping logs ([e3f5383](https://github.com/grindlemire/go-tui/commit/e3f53835851ea130f35f300dc820ce15bad5e45d))

## [0.9.0](https://github.com/grindlemire/go-tui/compare/v0.8.0...v0.9.0) (2026-03-22)


### Features

* reject @if/@for/@else/[@let](https://github.com/let) with diagnostic errors, remove TokenAtLet and parseLet ([0dae91b](https://github.com/grindlemire/go-tui/commit/0dae91b0047a341ed122cd5992c0fed1b9adac02))
* remove @-prefixed control flow from tree-sitter grammar ([6bd0d78](https://github.com/grindlemire/go-tui/commit/6bd0d78c099f6cf5bf784e15ece7175b26cf40c2))
* remove legacy @-prefixed control flow syntax ([cd711cb](https://github.com/grindlemire/go-tui/commit/cd711cb2e75b325442a59267b283933f6feedc21))


### Bug Fixes

* panic on unreachable LetBinding variant, suppress duplicate error for [@let](https://github.com/let) ([4dc7ad3](https://github.com/grindlemire/go-tui/commit/4dc7ad31e0e0e0158506fda00f62f49708a28bdd))

## [0.8.0](https://github.com/grindlemire/go-tui/compare/v0.7.0...v0.8.0) (2026-03-18)


### Features

* add Events, DispatchEvents, Step methods for manual event loops ([419fcd6](https://github.com/grindlemire/go-tui/commit/419fcd6b42c9c18e3e724ccfc3d178c80412f055))
* add Open lifecycle method, make Close idempotent ([1d2362a](https://github.com/grindlemire/go-tui/commit/1d2362adc189af097bbb83be51ed78321ce63321))
* add UpdateEvent type for queued closures ([3a85fce](https://github.com/grindlemire/go-tui/commit/3a85fce643e6fb013e219a94c5051d810c7d173b))
* manual event loop API (Open/Events/Dispatch/Render/Close) ([973bf89](https://github.com/grindlemire/go-tui/commit/973bf892dcc5780b6d24620f2eb2c19228cc0e3a))
* split event channels with input priority fan-in ([60730a0](https://github.com/grindlemire/go-tui/commit/60730a0d59333cae09bb42293fb1f5ee7c392918))


### Bug Fixes

* address code review findings ([142f8bb](https://github.com/grindlemire/go-tui/commit/142f8bb6c87588293e62131624365f125ab65279))
* auto-scroll message feed to bottom ([7495525](https://github.com/grindlemire/go-tui/commit/7495525a494ae2835cb554fc989bd3c51492f512))
* eliminate TOCTOU race in Run/Open and normalize struct alignment ([940ddd9](https://github.com/grindlemire/go-tui/commit/940ddd90b342a7824efb24969100d73d9616846d))
* stop background goroutines on NewApp error path ([10d630d](https://github.com/grindlemire/go-tui/commit/10d630db63a751df5a58c115375e7fea52b7c237))
* use proper sticky-scroll pattern for message feed ([bf79eb7](https://github.com/grindlemire/go-tui/commit/bf79eb7987c185180c09f87e90746b56121d7e88))

## [0.7.0](https://github.com/grindlemire/go-tui/compare/v0.6.1...v0.7.0) (2026-03-15)


### Features

* add ApplyDim and FillBlank methods to Buffer ([04ad1ef](https://github.com/grindlemire/go-tui/commit/04ad1eff5aef8d4e1c18f85987afc135dc09d2f3))
* add auto focus highlighting, Tab cycling, key suppression, and bottom-sheet modal example ([30fa430](https://github.com/grindlemire/go-tui/commit/30fa430c16df7e82a2bd00a93b410bce07ca09f3))
* add Modal component with Render, KeyListener, and MouseListener ([6f414da](https://github.com/grindlemire/go-tui/commit/6f414da63b7d9a70269ed1070fccb9c8c4e32785))
* add modal support ([87d3d78](https://github.com/grindlemire/go-tui/commit/87d3d78c40f2741bbc127915094fc068b1f3eadc))
* add modal support to code generator ([56a62fb](https://github.com/grindlemire/go-tui/commit/56a62fbf1a1280e253c6a9cb239ef26f952afb6b))
* add modal to analyzer known tags and attributes ([c4b27eb](https://github.com/grindlemire/go-tui/commit/c4b27eb66adfbc75d8d3c3364c1e059d1a0d2917))
* add overlay flag to Element for modal support ([14976ed](https://github.com/grindlemire/go-tui/commit/14976ed7c84b618470feef95a6dccd4b36fded2b))
* add overlay registration system to App ([5519970](https://github.com/grindlemire/go-tui/commit/55199705df45abfe49c00e969ab81e99c8cf1e69))
* add overlay render pass and focus scoping for modal support ([1fc431e](https://github.com/grindlemire/go-tui/commit/1fc431e4ca82579215fac0e10225eb7fff93da17))
* add reset confirmation modal to color mixer example ([aff5dc4](https://github.com/grindlemire/go-tui/commit/aff5dc4af06087c2941770ed1faf975b5d23420c))
* **lsp:** add modal element definition to schema ([4c626e4](https://github.com/grindlemire/go-tui/commit/4c626e4149cee68cee130f723a2bb0b387ae50b5))


### Bug Fixes

* add Enter key handler to select focused modal button in color mixer ([e0de594](https://github.com/grindlemire/go-tui/commit/e0de594df32c7c3ed09da5290a79a2b1dc167262))
* address review feedback for modal PR ([10addd5](https://github.com/grindlemire/go-tui/commit/10addd594cd2abd6148268339e27b44e746273a3))
* ensure modal content children have opaque background by default ([680c6af](https://github.com/grindlemire/go-tui/commit/680c6af0883597cc28004a5dfb0832293ea48d07))
* guard modals in inline mode, fix mouse coords, and resolve dispatch validation ([e898946](https://github.com/grindlemire/go-tui/commit/e89894651f972d7855439d4aac0b521665aec741))
* repair RenderFull overlay handling and modal inline mode key blocking ([5917456](https://github.com/grindlemire/go-tui/commit/59174561cf5ad73cad5cbb98f0150638b6ff283b))
* resolve render consistency, focus lifecycle, and test robustness ([e300abf](https://github.com/grindlemire/go-tui/commit/e300abfb1e10038bb35677c34d96c78ffa42437f))
* **suspend:** prevent Kitty keyboard response from leaking as typed input after Ctrl+Z resume ([2016f29](https://github.com/grindlemire/go-tui/commit/2016f29115425fc650d8fc05b7d54a03a9380d13))

## [0.6.1](https://github.com/grindlemire/go-tui/compare/v0.6.0...v0.6.1) (2026-03-15)


### Bug Fixes

* **lsp:** highlight Go keywords in expression attribute values ([b147ae7](https://github.com/grindlemire/go-tui/commit/b147ae799f44c0a7743df093599b85dfb68d8614))

## [0.6.0](https://github.com/grindlemire/go-tui/compare/v0.5.0...v0.6.0) (2026-03-15)


### Features

* add KeyMatcher interface with On/OnStop/OnFocused constructors ([99c8002](https://github.com/grindlemire/go-tui/commit/99c8002d0eef138281becdb924c558a3cc7c2002))
* add Kitty keyboard negotiation to Terminal interface ([044e266](https://github.com/grindlemire/go-tui/commit/044e266af46f56345ba033a4213a8c7e287c78a5))
* add Kitty keyboard protocol escape sequences ([56e6938](https://github.com/grindlemire/go-tui/commit/56e693807502dfad2222759e22a9caeead498409))
* add Kitty keyboard protocol for Ctrl+H/I/M disambiguation ([591fbde](https://github.com/grindlemire/go-tui/commit/591fbde4cffe63a7c765434138cbc649076847cc))
* add KittyKeyboard capability and WithLegacyKeyboard option ([19cb4fc](https://github.com/grindlemire/go-tui/commit/19cb4fcc70be5d19e71f4d04bf668d3ac62ed6b5))
* implement Kitty keyboard protocol negotiation in ANSITerminal ([cac462e](https://github.com/grindlemire/go-tui/commit/cac462e2b262374c85fc298326f4d47f86e52411))
* include kitty-keyboard in capabilities string ([4a55000](https://github.com/grindlemire/go-tui/commit/4a550005463195d96528206e9900bd9860c57ae4))
* parse Kitty keyboard protocol CSI u sequences ([4f48121](https://github.com/grindlemire/go-tui/commit/4f48121211c9c5d565526e713a19c0f79ff9d742))
* wire Kitty keyboard negotiation into app lifecycle ([5d7aaaa](https://github.com/grindlemire/go-tui/commit/5d7aaaa208c37db6b273f09e9aeb8ce7525c4c3c))


### Bug Fixes

* add trailing newline to reader_types.go and document Rune(0) behavior ([fb52f4e](https://github.com/grindlemire/go-tui/commit/fb52f4e107dd92fa34f19b374129186722e3c9db))
* address code review findings for Kitty keyboard and terminal flush ([cdddab8](https://github.com/grindlemire/go-tui/commit/cdddab89007da25f2ac07e26a161144b5b83434a))
* address code review findings for Kitty keyboard implementation ([1f9def9](https://github.com/grindlemire/go-tui/commit/1f9def97f96e8c608919802d611828acb8c26907))
* handle Kitty keyboard on suspend/resume and clean up minor issues ([6f54fa4](https://github.com/grindlemire/go-tui/commit/6f54fa4c46b688c34da5a1435bb98337118506d6))
* make KeyCtrlH/I/M aliases for KeyBackspace/Tab/Enter ([4decb9a](https://github.com/grindlemire/go-tui/commit/4decb9a8d987ac9b427b434e8ce0b904cea58513))
* remove flag-2+ dead code from kittySpecialKeys and harden Kitty negotiation ([1dd9c19](https://github.com/grindlemire/go-tui/commit/1dd9c19b51443462d6d6c201caaeed077729378e))
* struct alignment via gofmt, smarter response terminator detection ([40551c0](https://github.com/grindlemire/go-tui/commit/40551c0c993de345b0a004fbca121aeb681e1b6b))

## [0.5.0](https://github.com/grindlemire/go-tui/compare/v0.4.0...v0.5.0) (2026-03-14)


### Features

* drop @ prefix from control flow keywords ([470db04](https://github.com/grindlemire/go-tui/commit/470db04495a876cf4b4f52e9527151572353545b))
* extend LetBinding AST for component call and expression RHS ([888a780](https://github.com/grindlemire/go-tui/commit/888a78008d3c5096cc0876ee79e9396ed0d4d229))
* formatter emits bare if/for/else and := instead of @-prefixed syntax ([8aefebe](https://github.com/grindlemire/go-tui/commit/8aefebe396583b48bf0b11a362ad7c6cd74487e2))
* generator handles component call and expression RHS in LetBinding ([9d68ba0](https://github.com/grindlemire/go-tui/commit/9d68ba0b348f1952be6262d5470a79ac87ab62a7))
* lexer and parser accept bare if/for/else, remove TokenAtIf/TokenAtFor/TokenAtElse ([8560391](https://github.com/grindlemire/go-tui/commit/85603914c6a8ceedee4faad7d6bc7e57edbd1bd7))
* LSP uses bare keyword syntax in completions, hover, and semantic tokens ([5cfc3de](https://github.com/grindlemire/go-tui/commit/5cfc3de4eef931db4d4e60b00be8ade71b13581d))
* parse bare if/for/else and := / var bindings with element/component RHS ([55a8047](https://github.com/grindlemire/go-tui/commit/55a8047e0b79609ccce8705345a7784013465b02))
* support multi-line component call arguments ([761cd7f](https://github.com/grindlemire/go-tui/commit/761cd7f110d7e0cb649d36710ca67718b41e457c))
* TextMate grammar highlights bare if/for/else and := bindings ([27d98e7](https://github.com/grindlemire/go-tui/commit/27d98e733a7602af65fa104221733c91b5d1d169))
* tree-sitter grammar supports bare if/for/else and := bindings ([80764b3](https://github.com/grindlemire/go-tui/commit/80764b350dfb3a6b06a7c59965781182252b9460))


### Bug Fixes

* analyzer handles component call and expression RHS in LetBinding ([3d86741](https://github.com/grindlemire/go-tui/commit/3d8674108868cf0c05173e77df1e6bdd6c88f755))
* **lsp:** drop @ prefix from offset calculations in LSP providers ([feecc14](https://github.com/grindlemire/go-tui/commit/feecc1491cd54c0ffdabadf0def6233e20cd18f5))
* **lsp:** prevent findElseKeyword false-positive on else inside strings ([394494e](https://github.com/grindlemire/go-tui/commit/394494ed848a16979423083edad963c380e71c9c))
* parseGoStatement stops at closing brace at depth 0 ([e83980f](https://github.com/grindlemire/go-tui/commit/e83980f40ac0ebd623b9ea920197168d2522ca67))
* **parser:** prevent speculative parsing from leaking errors and comments ([62caa56](https://github.com/grindlemire/go-tui/commit/62caa569d25c9f4b238358aeeb227fbd9b098a85))
* **parser:** use rune-aware iteration for UTF-8 correctness ([f542956](https://github.com/grindlemire/go-tui/commit/f54295637cabc6e2abc90cd417ff2bae77f27082))
* replace HTML-style comments with Go-style comments in docs code blocks ([1c287fe](https://github.com/grindlemire/go-tui/commit/1c287fe6c201be3295dc6baf7eee0b9d3d8f9d12))

## [0.4.0](https://github.com/grindlemire/go-tui/compare/v0.3.1...v0.4.0) (2026-03-13)


### Features

* **examples:** add directory tree data model and flatten logic ([02912e2](https://github.com/grindlemire/go-tui/commit/02912e23e3b32fd7b943a7a923272a1264ca41eb))
* **examples:** add directory tree example ([186b2d0](https://github.com/grindlemire/go-tui/commit/186b2d0d934b54bb20fda8ef8d9d3904174cfb8f))
* **examples:** add directory tree example skeleton ([639da53](https://github.com/grindlemire/go-tui/commit/639da53217172ebcfa9f25ac72010dc7b2e818be))
* **examples:** add directory tree keyboard navigation ([1bdae4e](https://github.com/grindlemire/go-tui/commit/1bdae4ee1adb027508d790f6ebde905e6bb17b8f))
* **examples:** add directory tree render and complete example ([480566e](https://github.com/grindlemire/go-tui/commit/480566e5fb2c281e6912fe7c853af76338a0103a))
* **examples:** add scrolling, overflow clipping, and ancestor path highlighting ([ee46208](https://github.com/grindlemire/go-tui/commit/ee46208ea46bf57d54ce444073a3286a4bbe7e11))
* **examples:** replace hardcoded tree with random generator ([b5aaa4b](https://github.com/grindlemire/go-tui/commit/b5aaa4b88f982231677779b012777c78c7d3b421))
* **examples:** show selected node path at top of directory tree ([4d7f0ee](https://github.com/grindlemire/go-tui/commit/4d7f0ee4f30d6ec0982cd2771ca2d5ed177a836c))


### Bug Fixes

* **docs:** move Google Fonts from CSS [@import](https://github.com/import) to HTML link tag ([e0ee0d0](https://github.com/grindlemire/go-tui/commit/e0ee0d0e8f9d7988875da85df4ae6e1d5d6a8670))
* **docs:** move playwright to devDependencies and fix tview widget count ([269e858](https://github.com/grindlemire/go-tui/commit/269e8588fffafff8b3906af4fdf5fdfdef07eab0))
* **docs:** remove unused useRef import ([7847f15](https://github.com/grindlemire/go-tui/commit/7847f157275d34c6827937d73e67bbcf02274bcb))
* **examples:** remove overflow-hidden that bypassed scrollable rendering ([39ea661](https://github.com/grindlemire/go-tui/commit/39ea6617256cfe8c32beee508cbd97f1af9f1fe2))
* **examples:** use state-driven scrollOffset for directory tree scrolling ([a272a91](https://github.com/grindlemire/go-tui/commit/a272a912b0e115f721a6684535ecc5eb782623d1))
* **lsp:** use character offsets for semantic tokens ([b7bcf07](https://github.com/grindlemire/go-tui/commit/b7bcf07653120b43d3b781bee7a3f2e4bdf4bc9a))
* **lsp:** use character offsets instead of byte offsets for semantic tokens ([1791cf6](https://github.com/grindlemire/go-tui/commit/1791cf6d757e891dce01719ef4ef75c0c3874e73))

## [0.3.1](https://github.com/grindlemire/go-tui/compare/v0.3.0...v0.3.1) (2026-03-09)


### Bug Fixes

* move SIGWINCH handling from reader to App and default to blocking input ([3b2879a](https://github.com/grindlemire/go-tui/commit/3b2879ab99fbbfae8fda2360df1041bef2db6035))

## [0.3.0](https://github.com/grindlemire/go-tui/compare/v0.2.0...v0.3.0) (2026-03-09)


### Features

* add autoFocus attribute and make focusColor optional for Input and TextArea ([184af94](https://github.com/grindlemire/go-tui/commit/184af94620528453fd9615d39959951d0d7fdd9a))
* add borderGradient, focusGradient, and backspace scroll fix for Input and TextArea ([d1437fd](https://github.com/grindlemire/go-tui/commit/d1437fd0a85cece58ad68bc7b7071c64861255ca))
* add FocusRequired flag to KeyPattern and focus-gated binding helpers ([5d97ee1](https://github.com/grindlemire/go-tui/commit/5d97ee153fabb842690fcb72c67e2c6fd768220a))
* add IsFocused query method to focusManager ([681ffb3](https://github.com/grindlemire/go-tui/commit/681ffb3e3fd450a1b56be15b8c38e0d440e22682))
* add reactive value binding and focusColor for Input and TextArea ([35d1679](https://github.com/grindlemire/go-tui/commit/35d167916c9058f6f89d822d3b1d1cfc127a32f7))
* implement focus-gated dispatch in dispatch table ([0ea047a](https://github.com/grindlemire/go-tui/commit/0ea047ad18050e78f966a00c955134e9875875b2))
* separate tabStop from focusable to fix Tab navigation ([4d6189d](https://github.com/grindlemire/go-tui/commit/4d6189d96d83ed1cc19c4cb84668c0283a75bf00))
* wire Input component into focus system with focus-gated key bindings ([d2689fa](https://github.com/grindlemire/go-tui/commit/d2689fa29436bee0fb06aecc871e0f0d7dc2502c))
* wire TextArea component into focus system with focus-gated key bindings ([af57b8f](https://github.com/grindlemire/go-tui/commit/af57b8fdc92008c6eb19a780567d7739cd3a253c))


### Bug Fixes

* add Tab/Shift+Tab focus navigation to elements example ([7f77e38](https://github.com/grindlemire/go-tui/commit/7f77e38f410b34e36671a2345ab932e41ac3bdc1))
* exclude focus-gated entries from dispatch table conflict validation ([1545aa6](https://github.com/grindlemire/go-tui/commit/1545aa6c51ecf40f7c6ff69fa35a1d80656f5901))
* make ContainsPoint account for scroll offset in scrollable containers ([23be70e](https://github.com/grindlemire/go-tui/commit/23be70eebcd4ba8a04961fb361f6c908ec540717))

## [0.2.0](https://github.com/grindlemire/go-tui/compare/v0.1.2...v0.2.0) (2026-03-07)


### Features

* add default Ctrl+Z suspend fallback in key dispatch ([8f61d4e](https://github.com/grindlemire/go-tui/commit/8f61d4e161a5815f3ce7d267d20fb6be256eea65))
* add FlexWrap and AlignContent types with public API and Tailwind classes ([28eb41f](https://github.com/grindlemire/go-tui/commit/28eb41fedb83b8ac7f53a72399b03fa31524a11f))
* add no-op suspend stubs for Windows ([0484abf](https://github.com/grindlemire/go-tui/commit/0484abff025240f9190e50912fdc2912f07d2dcb))
* add OnChange watcher for reactive state effects ([623b93f](https://github.com/grindlemire/go-tui/commit/623b93fbb7a0e21db4bbdfa8d5667dd8cdd5c652))
* add OnKeyMod helper, fix flex-wrap align-content, update docs ([dbcce85](https://github.com/grindlemire/go-tui/commit/dbcce85edfed32c2780be63fd4de5d9a16f74c8b))
* add onSuspend/onResume fields to App struct ([f8be019](https://github.com/grindlemire/go-tui/commit/f8be019ef4ece3f4f4cbebfda5ef30cd0cb765b8))
* add WithOnSuspend and WithOnResume app options ([610b3ad](https://github.com/grindlemire/go-tui/commit/610b3ad76d69c927834eea455d42e4d541013ce1))
* implement suspend/resume terminal state management ([699e343](https://github.com/grindlemire/go-tui/commit/699e3435a3880e77cc2e2e59ada4ca4d1fab415c))
* **layout:** implement flex-wrap line breaking, per-line layout, and auto cross-axis sizing ([6932cd5](https://github.com/grindlemire/go-tui/commit/6932cd57dfd31eafa8b420655d05c5add0903ed3))
* register SIGTSTP signal handler in app event loop ([8e7163e](https://github.com/grindlemire/go-tui/commit/8e7163ee52d6321c6a40cbf70244ead99e007d12))


### Bug Fixes

* bake inline widget to scrollback before suspend ([f076cb3](https://github.com/grindlemire/go-tui/commit/f076cb369e41b3a97cb06af234feb393951902dc))
* clear widget area on inline suspend instead of baking duplicate ([a1ee2c2](https://github.com/grindlemire/go-tui/commit/a1ee2c247e50f6745f719aa72781195a2ec4df48))
* handle inline mode suspend/resume without corrupting scrollback ([03c1aae](https://github.com/grindlemire/go-tui/commit/03c1aae3574974b5b3851c9a02f0a64705bea6a4))
* prevent stack overflow from circular state dependencies ([f1c9ba5](https://github.com/grindlemire/go-tui/commit/f1c9ba58ec1e81b851f307319f6af8b83cbdc0b4))
* re-register SIGTSTP signal handler after resume ([554fded](https://github.com/grindlemire/go-tui/commit/554fdedbad37f379f92da0547bb454d94c79ec6b))
* replace sleep-based test sync with done channel, document FlexGrow heuristic ([dcc95f0](https://github.com/grindlemire/go-tui/commit/dcc95f05ad8d0aef43f5eac90b4b53af8debc1df))
* resolve three suspend/resume issues ([7e5b061](https://github.com/grindlemire/go-tui/commit/7e5b0614c918c873187891d0433ddafe17800185))

## [0.1.2](https://github.com/grindlemire/go-tui/compare/v0.1.1...v0.1.2) (2026-03-04)


### Bug Fixes

* use runtime/debug to report version from go install ([bd586f8](https://github.com/grindlemire/go-tui/commit/bd586f84e8c465e20f4d36b980721d077c27ec4b))

## [0.1.1](https://github.com/grindlemire/go-tui/compare/v0.1.0...v0.1.1) (2026-03-04)


### Bug Fixes

* update README badges, examples, and VS Code extension docs ([8c9fdf4](https://github.com/grindlemire/go-tui/commit/8c9fdf444d7bdbcd29633b7b7c1aedb249cd62d5))

## 0.1.0 (2026-03-03)


### Features

* add 1-character minimum spacing between table columns ([0036fc0](https://github.com/grindlemire/go-tui/commit/0036fc01828fb4983b96ce554766b8d488e2237e))
* add ANSI-aware byte scanner for styled streaming ([37c51d3](https://github.com/grindlemire/go-tui/commit/37c51d36a06b35717167baab960bcecc10f4e78f))
* add App.StreamAbove() with PrintAbove coordination ([09f6eac](https://github.com/grindlemire/go-tui/commit/09f6eacc14675fcef1d804c7dca024c8aae68e75))
* add bufferRowToANSI for rendering buffer rows to ANSI strings ([a2c1db9](https://github.com/grindlemire/go-tui/commit/a2c1db96dbb11b34ae820c89f376b14276f335cb))
* add ClickBinding type for ref-based mouse handling ([a18bf5b](https://github.com/grindlemire/go-tui/commit/a18bf5b0b90f47e4e807845db9e9c5d41031fb44))
* add Element.Component() public getter ([2c489a5](https://github.com/grindlemire/go-tui/commit/2c489a501e4d76439a8691221df91137c0a9b5d9))
* add ElementTag field for table layout dispatch ([51daa21](https://github.com/grindlemire/go-tui/commit/51daa21251c0cb845300f8b06cce6fbadc574176))
* add generic NewChannelWatcher helper ([52f3f86](https://github.com/grindlemire/go-tui/commit/52f3f86082ef2285a04f5b684ada9ceabf92edc7))
* add HandleClicks helper for automatic ref hit testing ([ba91e9d](https://github.com/grindlemire/go-tui/commit/ba91e9d9d6b79c06cb16eb6198ad635450bfa5fa))
* add HeightForWidth to Layoutable for text wrap height calculation ([5edcd48](https://github.com/grindlemire/go-tui/commit/5edcd484f554040cf5cd56feefef6a11adb3a752))
* add inlineStreamWriter and nopStreamWriter types ([e180a59](https://github.com/grindlemire/go-tui/commit/e180a594f8982cdaf7fad109ca760be53e6784b6))
* add MountPersistent to prevent component sweep when hidden ([0b8ce63](https://github.com/grindlemire/go-tui/commit/0b8ce6345c1edaa5a97fc1fa7b75e074bc2f7c38))
* add nowrap and wrap tailwind classes ([43e4b5a](https://github.com/grindlemire/go-tui/commit/43e4b5ae72b30bdf724627b33692a6c798a09939))
* add partial line tracking and appendBytes to inlineSession ([cb55fd3](https://github.com/grindlemire/go-tui/commit/cb55fd35a4a2c7ab937c75e782541561e9540bf9))
* add Print, Sprint, and Fprint for single-frame rendering ([c7b4660](https://github.com/grindlemire/go-tui/commit/c7b46603647f8147cf4f4d02e30807eb33185a2e))
* add Print, Sprint, and Fprint for single-frame rendering ([6c580b5](https://github.com/grindlemire/go-tui/commit/6c580b58ad1fb8df01cf8cb0cf3f428dbef78ae9))
* add PrintAboveElement and QueuePrintAboveElement for inline mode ([60e23c5](https://github.com/grindlemire/go-tui/commit/60e23c520dfae9131bb9ba63f78eccf06b537480))
* add renderElementToBuffer for standalone element rendering ([2899fc6](https://github.com/grindlemire/go-tui/commit/2899fc6bdcf838d51f2515ca3cbdfb00ca3dfd4d))
* add single-frame print example (examples/19-print) ([8dd1aef](https://github.com/grindlemire/go-tui/commit/8dd1aeffed882f9762b4ce3449190e46776de0f0))
* add StreamWriter.WriteElement for mid-stream element insertion ([fc7e4ae](https://github.com/grindlemire/go-tui/commit/fc7e4ae92c38c50d111d693297daf38504a769f5))
* add tree walk to collect watchers from WatcherProvider components ([2abe659](https://github.com/grindlemire/go-tui/commit/2abe659c861332e3f8cc3772c2f1e4b3d88244f6))
* add WatcherProvider interface for component-level watchers ([4ceac54](https://github.com/grindlemire/go-tui/commit/4ceac54f40f893b1531a4f83adc69a69a409974f))
* add WithWrap option for text wrapping (default enabled) ([8efdf97](https://github.com/grindlemire/go-tui/commit/8efdf97a0d5502787682847140d3120846160d9a))
* add word-wrap text function with mid-word fallback ([2632126](https://github.com/grindlemire/go-tui/commit/263212657d99ff948601d0d81ef6ddca184f707d))
* add Wrap() public getter to Element ([448c1cd](https://github.com/grindlemire/go-tui/commit/448c1cde54dfdd054ae8a00d2e4efedefa2c4809))
* **ai-chat:** add ChatApp root component with streaming ([ff392d4](https://github.com/grindlemire/go-tui/commit/ff392d416d07cce98d95c41035cb883730c5015d))
* **ai-chat:** add fake provider for demo/testing ([06c006c](https://github.com/grindlemire/go-tui/commit/06c006cfd4e302ecc62aeff5783ea91c13b3d44b))
* **ai-chat:** add Header component ([038fad6](https://github.com/grindlemire/go-tui/commit/038fad6ef622b690f81101af6aa836f583675903))
* **ai-chat:** add HelpOverlay component ([8168250](https://github.com/grindlemire/go-tui/commit/8168250f057cd9b307b6ead34fe3d408ce5be2d2))
* **ai-chat:** add Message component with copy/retry actions ([4e95b3f](https://github.com/grindlemire/go-tui/commit/4e95b3f156c1e5b260cd5f1af3d32259d08a069e))
* **ai-chat:** add MessageList component with vim navigation ([e7e7950](https://github.com/grindlemire/go-tui/commit/e7e7950e1827ca81d66e92a5db2ff7d7bae18ee3))
* **ai-chat:** add provider abstraction for OpenAI/Anthropic/Ollama ([2e0cac0](https://github.com/grindlemire/go-tui/commit/2e0cac0dd7491d2cd4179ee19f24901a70fce4b9))
* **ai-chat:** add Settings screen components ([d27cf59](https://github.com/grindlemire/go-tui/commit/d27cf59fe6e96a410db9485ea41ed028e98a01f7))
* **ai-chat:** add settings screen entry point ([3e86430](https://github.com/grindlemire/go-tui/commit/3e86430ce37df748a87f9683f565ff377352917a))
* **ai-chat:** add state types and AppState ([e87e9d3](https://github.com/grindlemire/go-tui/commit/e87e9d3576597753c07f0923b06dd1632a233d54))
* **ai-chat:** migrate textarea to GSX element with persistent mount ([f83b016](https://github.com/grindlemire/go-tui/commit/f83b0169e1afe734391909703df911c2144f305e))
* **ai-chat:** wire settings screen to main app ([66fe6ad](https://github.com/grindlemire/go-tui/commit/66fe6ad8872c35dfd0cda5ec6101a5845ae6b1f7))
* **ai-chat:** wire up ChatApp with provider detection ([d587eb5](https://github.com/grindlemire/go-tui/commit/d587eb52af8326ed1c21cfbcd3d3141076267548))
* **analyzer:** accept textarea as valid element tag ([0fdb81a](https://github.com/grindlemire/go-tui/commit/0fdb81abf08a2b313db6c5970a8ed7b7a75396f4))
* auto-scroll wrapped text with hidden scrollbar on overflow ([b3fa83c](https://github.com/grindlemire/go-tui/commit/b3fa83c0db75d03cbdbc7e2a63f6ce5c0d7af983))
* compute table intrinsic size from column widths and row heights ([b5b9634](https://github.com/grindlemire/go-tui/commit/b5b963464633edb69ed6dc1db5cb56a60b767371))
* **element:** add integration tests and update dashboard example ([326e7f9](https://github.com/grindlemire/go-tui/commit/326e7f91eb72b57556cba0ec21b74fb8c9952b9d))
* **element:** add onUpdate hook for pre-render callbacks ([551fbd2](https://github.com/grindlemire/go-tui/commit/551fbd2f233cff43677a6658b4ad8ef52abeeee1))
* **element:** implement Phase 1 - Layout interface and Element core ([a3211f4](https://github.com/grindlemire/go-tui/commit/a3211f4a05c6ec048e345c3f0352939c231043be))
* **examples:** scaffold ai-chat example ([c8912ef](https://github.com/grindlemire/go-tui/commit/c8912ef4aa4562312295182208afd19597cafda6))
* generate WithTag for table elements in codegen ([a854717](https://github.com/grindlemire/go-tui/commit/a854717a45054d331cddc8d94f74341ef676b435))
* **generator:** emit app.Mount + NewTextArea for textarea elements ([31c29d7](https://github.com/grindlemire/go-tui/commit/31c29d7031f9760c011170b460fd54be1b7ebba3))
* **generator:** emit MountPersistent for component elements ([2635176](https://github.com/grindlemire/go-tui/commit/26351768a47f326ecb8f9fa7b71516323b577218))
* implement table layout algorithm with auto column sizing ([5d1d76a](https://github.com/grindlemire/go-tui/commit/5d1d76af49d1cac71dfa016221223d33d970b8e9))
* integrate WatcherProvider watchers into app lifecycle ([cbc588f](https://github.com/grindlemire/go-tui/commit/cbc588f866b7432700eca499e1e5b6462647c87d))
* register tr, td, th elements in schema and analyzer ([20a41d9](https://github.com/grindlemire/go-tui/commit/20a41d98b9dc1970b1900f45294e8b286bdccbb9))
* render th elements with bold text by default ([c7a0849](https://github.com/grindlemire/go-tui/commit/c7a084945f86e8503370bb893ea496383c6fbec8))
* render wrapped text across multiple lines with per-line alignment ([12df559](https://github.com/grindlemire/go-tui/commit/12df5598b14c05baf254fd2b735b69e3891c9029))
* restore EventInspector with event tracking in interactive example ([ffc4774](https://github.com/grindlemire/go-tui/commit/ffc47741ebf6dc73f7509bc90c9891d73fce7d38))
* **schema:** add textarea element definition to LSP schema ([d141b42](https://github.com/grindlemire/go-tui/commit/d141b4277adccd5e158ae13e45089fe6c0641bdb))
* **tailwind:** add validation, similarity matching, and class registry (Phase 2) ([935db99](https://github.com/grindlemire/go-tui/commit/935db99d848fe89b6942a73a9afe925a68324e34))
* **tailwind:** expand class mappings with percentages, individual sides, and flex utilities (Phase 1) ([abe07ef](https://github.com/grindlemire/go-tui/commit/abe07eff7d46f6a2d05ac6b600ef97a2cd905018))
* **tailwind:** integrate class validation into analyzer and LSP diagnostics (Phase 3) ([3e1ac96](https://github.com/grindlemire/go-tui/commit/3e1ac966230947db6b01f121d397ab7c46972056))
* **tuigen:** add named element refs syntax (#Name) - Phase 1 ([91af36a](https://github.com/grindlemire/go-tui/commit/91af36ada865970907c43c44db0e9ac034fa0b8e))
* **tuigen:** add named element refs syntax (#Name) - Phase 1 ([6e7e33a](https://github.com/grindlemire/go-tui/commit/6e7e33ad27771326de5372f516738b9d92992291))
* **tuigen:** add named element refs syntax (#Name) - Phase 1 ([0e47fa4](https://github.com/grindlemire/go-tui/commit/0e47fa44551c465b740cea0f13d03fe746785d0f))
* **tuigen:** add state detection to analyzer - Phase 3 ([17fb834](https://github.com/grindlemire/go-tui/commit/17fb8349a5d8bf028fc62d885d18719a85bb021a))
* **tui:** implement App.Run(), SetRoot(), and element handlers - Phases 2 & 3 ([ba97b9a](https://github.com/grindlemire/go-tui/commit/ba97b9a7b82735e1d18070f790cd057bc834a518))
* **tui:** implement Batch() for coalescing state updates - Phase 2 ([cee1b25](https://github.com/grindlemire/go-tui/commit/cee1b25ddda813884b748a4f049e5b6da1340a41))
* **tui:** implement dirty tracking and watcher types - Phase 1 ([162a1c0](https://github.com/grindlemire/go-tui/commit/162a1c0bc289f3b9720dd5d99d457efd83ed0eee))
* **tui:** implement State[T] reactive type with bindings - Phase 1 ([99665a1](https://github.com/grindlemire/go-tui/commit/99665a143c2e8ca0dc393a29ca4eb0f8767249e3))
* two-pass layout for text wrapping height calculation ([afb032f](https://github.com/grindlemire/go-tui/commit/afb032f3e5435c331a7f9c015e563e7cd51118f9))


### Bug Fixes

* add explicit row height override for table tr elements ([0f50303](https://github.com/grindlemire/go-tui/commit/0f50303b24e5a7c047bc3aad32aae59277c8e3eb))
* **ai-chat:** fix help text, temperature labels, and copy handler ([3f10c71](https://github.com/grindlemire/go-tui/commit/3f10c71c830c24f3ee389d11800efbe5cec4a16a))
* **ai-chat:** remove broken go:generate directive ([7fa71b2](https://github.com/grindlemire/go-tui/commit/7fa71b27ef11149945766590e05d2f5dacebd94a))
* **ai-chat:** use explicit style attrs for dynamic styling ([52ce3e3](https://github.com/grindlemire/go-tui/commit/52ce3e39d5d2ec9144174cd33146b688e1728f7f))
* align panel heights in interactive example ([1953b39](https://github.com/grindlemire/go-tui/commit/1953b39295e702cad21b1ca604cb1cb344bc02f2))
* flaky tests, nil-ref panic, and fill test coverage gaps ([9c3491c](https://github.com/grindlemire/go-tui/commit/9c3491c8bbc8aa29b2bfe847385047a5737c28f8))
* guard against nil ref in HandleClicks ([2380913](https://github.com/grindlemire/go-tui/commit/238091376131c172900a0fac1bdfb31b0fd34609))
* move textElementWithOptions/skipTextChildren to generator_element.go per spec ([4f35a8f](https://github.com/grindlemire/go-tui/commit/4f35a8fa476ca173ba09550b73797ea0554afdfb))
* prevent header from shrinking with flexShrink={0} ([ab13198](https://github.com/grindlemire/go-tui/commit/ab131986faa8e2d1243238384a4ca63371de4162))
* recursive HeightForWidth for containers to propagate wrapped text height ([28c3486](https://github.com/grindlemire/go-tui/commit/28c34865922157777b8d6225d4952069c446761f))
* replace time.Sleep with channel draining in watcher tests ([994d0cb](https://github.com/grindlemire/go-tui/commit/994d0cb96bfbda8da942695ba5c062abfe1d46b5))
* restore previous currentApp in Run() for nested apps ([bf07b93](https://github.com/grindlemire/go-tui/commit/bf07b93575223b8d33b8dec3ee409317cc8b1713))
* settings screen as embedded component instead of separate app ([123cf98](https://github.com/grindlemire/go-tui/commit/123cf980516d45d899738aad91b6e82b95306fc9))
* use flexGrow for proper vertical distribution in interactive example ([cc10eb1](https://github.com/grindlemire/go-tui/commit/cc10eb10b1e6549aceca1f516b903abf93a4da0d))
