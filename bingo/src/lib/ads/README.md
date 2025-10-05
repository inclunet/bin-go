# Sistema de Anúncios do Jogo da Velha

Este sistema oferece controle completo sobre onde e como os anúncios do Google AdSense aparecem, garantindo que não interfiram na jogabilidade.

## 📋 Configuração Inicial

### 1. Configurar IDs do AdSense

Edite o arquivo `src/lib/ads/adConfig.js`:

```javascript
export const adConfig = {
  // Substitua pelo seu ID de cliente AdSense real
  client: 'ca-pub-1234567890123456',
  
  slots: {
    topBanner: '1234567890',      // Banner superior
    rectangle: '1234567891',      // Retângulo meio/lateral
    bottomBanner: '1234567892',   // Banner inferior
    sidebar: '1234567893'         // Barra lateral (desktop)
  },
  // ... resto da configuração
};
```

### 2. Incluir Script do AdSense

Adicione no `app.html` ou na página principal:

```html
<script async src="https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-XXXXXXX"
     crossorigin="anonymous"></script>
```

## 🛡️ Sistema de Proteção

### Níveis de Proteção

- **`strict`** (padrão): Proteção máxima, bloqueia anúncios automáticos
- **`medium`**: Proteção moderada, permite alguns anúncios controlados
- **`light`**: Proteção mínima, apenas elementos críticos

### Áreas Protegidas

O sistema automaticamente protege:
- ✅ Tabuleiro do jogo (`.tic-board`)
- ✅ Células do tabuleiro (`.tic-cell`)
- ✅ Placar (`.scoreboard-container`)
- ✅ Controles do jogo (`.game-controls`)
- ✅ Modais de ajuda
- ✅ Elementos interativos (botões, links)

## 📱 Responsividade

### Comportamento por Dispositivo

**Mobile (≤480px):**
- Apenas banners top/bottom
- Máximo 320x100px
- Carregamento atrasado

**Tablet (481-768px):**
- Banners + retângulos
- Máximo 728x150px

**Desktop (>768px):**
- Todos os formatos disponíveis
- Inclui sidebar (160x600px)

## 🎯 Posicionamento dos Anúncios

### Locais Estratégicos

1. **Banner Superior**: Aparece antes do placar (apenas desktop)
2. **Retângulo Central**: Após resultado do jogo
3. **Banner Inferior**: No final da página
4. **Sidebar**: Lateral direita (apenas desktop)

### Distâncias de Segurança

- **Do tabuleiro**: 100px mínimo
- **Do placar**: 50px mínimo  
- **Dos controles**: 80px mínimo

## ⚙️ Personalização

### Ajustar Configurações

```javascript
// Em adConfig.js
settings: {
  protectionLevel: 'strict',    // Nível de proteção
  loadDelay: 2000,             // Atraso de carregamento (ms)
  maxSizes: {                  // Tamanhos máximos por dispositivo
    mobile: { width: 320, height: 100 },
    tablet: { width: 728, height: 150 },
    desktop: { width: 728, height: 200 }
  }
}
```

### Desabilitar Anúncios Automáticos

O sistema automaticamente:
- ✅ Desabilita `enable_page_level_ads`
- ✅ Remove overlays automáticos
- ✅ Bloqueia anúncios âncora
- ✅ Adiciona meta tags de controle

## 🔧 Componentes

### AdManager.svelte
Gerencia anúncios individuais com controle fino de posicionamento e tamanho.

**Propriedades:**
- `placement`: 'top', 'bottom', 'sidebar', 'auto'
- `format`: 'banner', 'rectangle', 'square'
- `priority`: 'high', 'medium', 'low'
- `adSlot`: ID do slot do AdSense

### AdProtection.svelte
Aplica proteções CSS e JavaScript para evitar interferências.

**Propriedades:**
- `protection`: 'strict', 'medium', 'light'

## 📊 Monitoramento

### Logs de Debug

O sistema registra:
- Anúncios muito próximos de elementos do jogo
- Falhas de carregamento
- Conflitos de posicionamento

### Verificações Automáticas

- ✅ Distância mínima respeitada
- ✅ Tamanhos dentro dos limites
- ✅ Posicionamento adequado ao dispositivo
- ✅ Carregamento não-obstrutivo

## 🚀 Implementação

### Exemplo Básico

```svelte
<!-- Proteção geral -->
<AdProtection protection="strict" />

<!-- Anúncio controlado -->
<AdManager 
  placement="bottom" 
  format="banner" 
  adSlot={adConfig.slots.bottomBanner}
  priority="low"
/>
```

### Exemplo Avançado

```svelte
<!-- Com configuração personalizada -->
<AdManager 
  placement="auto" 
  format="rectangle" 
  adSlot="1234567890"
  priority="medium"
  responsive={true}
/>
```

## ✅ Boas Práticas

1. **Sempre use AdProtection** na página principal
2. **Configure distâncias adequadas** para o seu layout
3. **Teste em diferentes dispositivos** antes de ativar
4. **Use priority="low"** para anúncios não-críticos
5. **Monitore performance** e impacto na jogabilidade
6. **Respeite a experiência do usuário** - menos é mais

## 🐛 Solução de Problemas

### Anúncios não aparecem
- Verifique IDs do cliente/slot
- Confirme script do AdSense carregado
- Teste com proteção 'light'

### Interferência no jogo
- Aumente distâncias mínimas
- Use proteção 'strict'
- Verifique CSS personalizado

### Performance lenta
- Aumente `loadDelay`
- Use priority 'low'
- Implemente lazy loading