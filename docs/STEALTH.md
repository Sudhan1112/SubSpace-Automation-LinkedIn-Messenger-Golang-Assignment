# Stealth Strategy & Anti-Detection

This document outlines the techniques used in **SubSpace Automator** to simulate organic human behavior and avoid simple bot detection mechanisms. These strategies contribute to the **35% Anti-Detection Quality** evaluation criteria.

## 1. Human-Like Mouse Movements
**Technique**: Bézier Curve Trajectories
- **Why**: Bots typically move the mouse in straight lines (Point A to Point B) with constant speed. This is easily flagged.
- **Implementation**: We use the `go-rod/stealth` library which calculates Bézier curves for mouse movement. When clicking an element, the cursor doesn't teleport; it glides imperceptibly with varying velocity, simulating hand-eye coordination errors and corrections.

## 2. Randomized Time Delays
**Technique**: Stochastic Waiting
- **Why**: Completing actions instantly (e.g., clicking a button 0ms after page load) is a clear bot signal.
- **Implementation**:
  - **Type Latency**: Key presses have random intervals (e.g., 50ms - 150ms) to simulate typing speed variation.
  - **Think Time**: After page transitions or before clicks, a random "think time" (e.g., 2s - 5s) is injected.

## 3. Fingerprint Masking
**Technique**: Navigator Property Overrides
- **Why**: Headless browsers leak properties (e.g., `navigator.webdriver = true`) that websites use to identify automation.
- **Implementation**:
  - **`navigator.webdriver`**: Set to `false` or `undefined`.
  - **User-Agent**: Automatically rotates or sets a standard, non-headless User-Agent string.
  - **WebGL/Canvas**: Not explicitly obfuscated in this POC, but relying on Rod's default stealth evades basic checks.

## 4. Context Isolation
**Technique**: Incognito Contexts
- **Why**: Preventing cookie leakage and ensuring a fresh state for testing.
- **Implementation**: Each automation runs in an isolated browser context (Incognito mode). This mimics a clean user session and prevents cross-contamination of local storage data.

## 5. Limitations
- **Advanced Fingerprinting**: This POC does not fully mask advanced TLS fingerprinting or canvas fingerprinting used by high-end anti-bot systems (e.g., Cloudflare, Akamai).
- **IP Reputation**: The stealth layer cannot hide the IP address. Repeated aggressive requests from a data center IP will still trigger rate limits.
