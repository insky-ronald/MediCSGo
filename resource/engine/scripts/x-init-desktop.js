function PrepareDesktop() {
	var linkCss = document.createElement("link")
		linkCss.href = "/loadcss/app/css/main-mobile.css?pid=mobile";
		linkCss.rel = "stylesheet";
		linkCss.type = "text/css";
		
		document.head.appendChild(linkCss);
	
	var script = document.createElement('script');
		script.src = "/loadscript/app/scripts/main-mobile?pid=mobile";
		script.type = "text/javascript";
		script.charset = "utf-8";
		script.onload = function() {
			desktop = new MainPage({id:"home", mobile:true});
		};
	
		document.head.appendChild(script);
};
