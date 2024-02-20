window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">

  // 현재 URL에서 쿼리 파라미터를 파싱합니다.
  const queryParams = new URLSearchParams(window.location.search);
  // 'url' 쿼리 파라미터를 읽습니다. 제공되지 않는 경우, 기본값으로 '/swagger.json'을 사용합니다.
  const swaggerUrl = queryParams.get('url') || '/swagger.json';

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  window.ui = SwaggerUIBundle({
    url: swaggerUrl,
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  });

  //</editor-fold>
};
