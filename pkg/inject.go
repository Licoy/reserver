package pkg

func GetInjectScript(after string) string {
	return `
<!-- reServer monitor -->
<script type="text/javascript">
	function refreshCSS() {
		const sheets = [].slice.call(document.getElementsByTagName("link"));
		for (var i = 0; i < sheets.length; ++i) {
			const elem = sheets[i];
			if (elem.rel != "stylesheet") {
				continue
			}
			const href = elem.href;
			elem.href = href.substring(0, href.lastIndexOf(".css")) + ".css?r=" + (new Date().getTime())
		}
	}
	const protocol = window.location.protocol === 'http:' ? 'ws://' : 'wss://';
	const address = protocol + window.location.host + '/ws';
	const reServerSocket = new WebSocket(address);
	reServerSocket.onmessage = function (msg) {
		if (msg.data === 'reload') window.location.reload();
		else if (msg.data == 'css') refreshCSS();
	};
	reServerSocket.onopen = function () {
		console.log('ReServer reload enabled.');
	}
	reServerSocket.onclose = function () {
		console.error('ReServer reload closed.');
	}
	reServerSocket.onerror = function () {
		console.error('ReServer reload open failed.');
	}
</script>
` + after
}
