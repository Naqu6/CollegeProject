function callAPI(data, fn) {
	$.ajax("/api", {
		data: data,
		dataType: 'json',
		success: (data) => {
			if (fn) {
				fn(data);
			}
		}
	});
}

function splitCookieString() {
	var cookies = document.cookie.split(";");
	var cookieDict = {};
	cookies.forEach((cookie, _) => {
		var i = cookie.indexOf("=");
		cookieDict[cookie.slice(0, i).replace(" ", "")] = cookie.slice(i+1);
	});

	return cookieDict;
}