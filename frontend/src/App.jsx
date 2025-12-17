import { useState, useEffect } from 'react'
import { Play, Square, Terminal, Search, Lock, Activity } from 'lucide-react'

function App() {
  const [status, setStatus] = useState({ running: false, logs: [] })
  const [searchQuery, setSearchQuery] = useState("")
  const [creds, setCreds] = useState({ email: "", password: "" })
  const [showCreds, setShowCreds] = useState(false)

  // Polling for real-time updates
  useEffect(() => {
    const interval = setInterval(() => {
      fetchStatus()
    }, 2000)
    return () => clearInterval(interval)
  }, [])

  const fetchStatus = async () => {
    try {
      const res = await fetch('/api/status')
      const data = await res.json()
      setStatus(data)
    } catch (e) {
      console.error("Status fetch failed", e)
    }
  }

  const startAutomation = async () => {
    if (!searchQuery) return
    try {
      const res = await fetch('/api/start', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          query: searchQuery,
          email: creds.email,
          password: creds.password
        })
      })
      if (res.ok) fetchStatus()
    } catch (e) {
      console.error("Failed to start", e)
    }
  }

  const stopAutomation = async () => {
    try {
      const res = await fetch('/api/stop', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' }
      })
      if (res.ok) fetchStatus()
    } catch (e) {
      console.error("Failed to stop", e)
    }
  }

  // Professional Warm White Theme Palette
  const bgMain = "bg-[#fcfbf9]" // Very subtle warm white
  const bgCard = "bg-white"
  const inputBg = "bg-[#F5F5F4]" // Stone 100

  return (
    <div className={`min-h-screen w-full flex flex-col items-center justify-center p-6 md:p-12 ${bgMain} font-sans selection:bg-stone-200 selection:text-stone-800`}>

      {/* Main Container - Added subtle ring for depth */}
      <div className={`w-full max-w-7xl ${bgCard} rounded-[2rem] shadow-2xl shadow-stone-200/60 ring-1 ring-stone-900/5 p-8 md:p-12 min-h-[85vh] flex flex-col transition-all duration-300`}>

        {/* Header Section */}
        <div className="flex flex-col lg:flex-row items-center justify-between gap-8 mb-12">

          {/* Left: Custom Credentials Button */}
          <div className="relative z-30 w-full lg:w-auto">
            <button
              onClick={() => setShowCreds(!showCreds)}
              className={`group w-full lg:w-auto px-6 py-4 rounded-2xl text-sm font-semibold tracking-wide transition-all duration-300 ease-out border shadow-sm ${showCreds
                  ? "bg-stone-800 border-stone-800 text-stone-50 shadow-stone-800/20"
                  : "bg-white border-stone-200 text-stone-600 hover:border-stone-300 hover:shadow-md hover:-translate-y-0.5"
                }`}
            >
              <span className="flex items-center justify-center gap-2.5">
                <Lock size={16} className={showCreds ? "text-stone-300" : "text-stone-400 group-hover:text-stone-600 transition-colors"} />
                {showCreds ? "Hide Credentials" : "Custom Credentials(Optional Input)"}
              </span>
            </button>

            {showCreds && (
              <div className="absolute top-full left-0 mt-4 w-full lg:w-96 bg-white rounded-2xl shadow-2xl shadow-stone-400/20 ring-1 ring-stone-900/5 p-6 space-y-5 animate-scale-in origin-top-left z-50">
                <div className="space-y-2">
                  <label className="text-[11px] font-bold uppercase tracking-widest text-stone-400 ml-1">Email Address</label>
                  <input
                    type="email"
                    value={creds.email}
                    onChange={e => setCreds({ ...creds, email: e.target.value })}
                    className={`w-full px-4 py-3.5 ${inputBg} border-transparent rounded-xl text-sm font-medium outline-none focus:bg-white focus:ring-2 focus:ring-stone-800/10 focus:border-stone-300 transition-all text-stone-700 placeholder:text-stone-400`}
                    placeholder="name@example.com"
                  />
                </div>
                <div className="space-y-2">
                  <label className="text-[11px] font-bold uppercase tracking-widest text-stone-400 ml-1">Password</label>
                  <input
                    type="password"
                    value={creds.password}
                    onChange={e => setCreds({ ...creds, password: e.target.value })}
                    className={`w-full px-4 py-3.5 ${inputBg} border-transparent rounded-xl text-sm font-medium outline-none focus:bg-white focus:ring-2 focus:ring-stone-800/10 focus:border-stone-300 transition-all text-stone-700 placeholder:text-stone-400`}
                    placeholder="••••••••••••"
                  />
                </div>
              </div>
            )}
          </div>

          {/* Right: Search & Action */}
          <div className="flex flex-col sm:flex-row gap-4 w-full lg:w-auto items-center">
            {/* Search Bar */}
            <div className="relative group w-full sm:w-[28rem]">
              <div className="absolute inset-y-0 left-0 pl-5 flex items-center pointer-events-none">
                <Search size={18} className="text-stone-400 group-focus-within:text-stone-600 transition-colors" />
              </div>
              <input
                type="text"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                disabled={status.running}
                className={`w-full pl-12 pr-6 py-4 ${inputBg} border-transparent rounded-2xl text-stone-800 placeholder:text-stone-400 focus:bg-white focus:ring-2 focus:ring-stone-800/5 focus:border-stone-300 focus:shadow-lg focus:shadow-stone-200/50 transition-all outline-none font-medium`}
                placeholder="Search Bar"
              />
            </div>

            {/* Automation Button */}
            {status.running ? (
              <button
                onClick={stopAutomation}
                className="w-full sm:w-auto px-8 py-4 bg-red-50 hover:bg-red-100 text-red-600 border border-red-100 rounded-2xl font-bold transition-all flex items-center justify-center gap-2.5 shadow-sm hover:shadow-md active:scale-95 duration-200 whitespace-nowrap"
              >
                <Square size={20} fill="currentColor" className="opacity-90" />
                <span>Stop System</span>
              </button>
            ) : (
              <button
                onClick={startAutomation}
                className="w-full sm:w-auto px-8 py-4 bg-[#292524] hover:bg-[#1c1917] text-[#FAFAF9] rounded-2xl font-bold tracking-wide transition-all shadow-xl shadow-stone-900/10 hover:shadow-stone-900/20 hover:-translate-y-0.5 active:scale-95 active:translate-y-0 duration-200 flex items-center justify-center gap-2.5 whitespace-nowrap"
              >
                <Play size={20} fill="currentColor" className="opacity-90" />
                <span>Automation Button</span>
              </button>
            )}
          </div>
        </div>

        {/* Main Content: Logs Area */}
        <div className="flex-grow w-full relative rounded-2xl overflow-hidden bg-[#F5F5F4] ring-1 ring-inset ring-stone-900/5 flex flex-col items-center justify-center min-h-[500px] group transition-colors hover:bg-[#F0EFEE]">

          {/* Terminal Header Decoration (Mac-like dots) */}
          <div className="absolute top-5 left-5 flex gap-2 opacity-30 group-hover:opacity-50 transition-opacity">
            <div className="w-3 h-3 rounded-full bg-stone-400"></div>
            <div className="w-3 h-3 rounded-full bg-stone-400"></div>
            <div className="w-3 h-3 rounded-full bg-stone-400"></div>
          </div>

          {/* Center Text when empty */}
          {status.logs.length === 0 ? (
            <div className="text-center space-y-3 p-6 transition-all duration-500">
              <div className="inline-flex p-4 rounded-full bg-stone-200/50 mb-2">
                <Terminal size={32} className="text-stone-400" />
              </div>
              <h3 className="text-2xl font-bold text-stone-800 tracking-tight">Live System Logs</h3>
              <p className="text-stone-500 font-medium">System is ready. Start automation to view real-time operations.</p>
            </div>
          ) : (
            <div className="absolute inset-0 p-8 pt-16 overflow-y-auto font-mono text-[13px] leading-relaxed w-full text-left scrollbar-hide">
              <div className="flex flex-col gap-3 max-w-5xl mx-auto">
                {status.logs.map((log, i) => (
                  <div key={i} className="flex gap-4 items-start animate-fade-in group/log p-3 rounded-lg hover:bg-white/60 transition-colors border border-transparent hover:border-stone-200/50">
                    <span className="flex-shrink-0 opacity-40 text-xs font-bold tracking-wider pt-1 select-none text-stone-500">
                      {new Date().toLocaleTimeString([], { hour12: false, hour: '2-digit', minute: '2-digit', second: '2-digit' })}
                    </span>
                    <span className="font-medium text-stone-700 break-words">{log}</span>
                  </div>
                ))}
                {/* Pulsing cursor at the end */}
                <div className="ml-[5.5rem] w-2 h-4 bg-stone-400 animate-pulse mt-1"></div>
              </div>
              <div className="h-10"></div>
            </div>
          )}
        </div>

      </div>
    </div>
  )
}

export default App