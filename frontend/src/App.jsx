import { useState, useEffect } from 'react'
import { Activity, Play, Square, Database, Search, User, MessageSquare, Sun, Moon, Shield, Lock, Users } from 'lucide-react'
import axios from 'axios'

function App() {
  const [darkMode, setDarkMode] = useState(true)
  const [status, setStatus] = useState({ running: false, logs: [] })
  const [activities, setActivities] = useState([])

  const [searchQuery, setSearchQuery] = useState("")
  const [creds, setCreds] = useState({ email: "", password: "" })
  const [showCreds, setShowCreds] = useState(false)

  // Polling for real-time updates
  useEffect(() => {
    const interval = setInterval(() => {
      fetchStatus()
      fetchData() // Always fetch data to keep stats updated
    }, 2000)
    return () => clearInterval(interval)
  }, [])

  const fetchStatus = async () => {
    try {
      const res = await axios.get('/api/status')
      setStatus(res.data)
    } catch (e) {
      console.error("Status fetch failed", e)
    }
  }

  const fetchData = async () => {
    try {
      const res = await axios.get('/api/data')
      setActivities(res.data)
    } catch (e) {
      console.error("Data fetch failed", e)
    }
  }

  const startAutomation = async () => {
    if (!searchQuery) return alert("Please enter a job title or keyword.")
    try {
      await axios.post('/api/start', {
        query: searchQuery,
        email: creds.email,
        password: creds.password
      })
      fetchStatus()
    } catch (e) {
      alert("Failed to start automation. Check backend.")
    }
  }

  const stopAutomation = async () => {
    try {
      await axios.post('/api/stop')
      fetchStatus()
    } catch (e) {
      alert("Failed to stop.")
    }
  }

  // Calculated Stats
  const profilesFound = activities.filter(a => a.Action === 'SEARCH_FOUND').length
  const requestsSent = activities.filter(a => a.Action === 'CONNECT').length
  const messagesSent = activities.filter(a => a.Action === 'MESSAGE').length

  const themeClass = darkMode ? "bg-slate-900 text-white" : "bg-gray-50 text-gray-900"
  const cardClass = darkMode ? "bg-slate-800/50 border-slate-700 hover:border-slate-600" : "bg-white border-gray-200 hover:border-blue-300"
  const inputClass = darkMode ? "bg-slate-900 border-slate-700 text-white focus:border-blue-500" : "bg-white border-gray-300 text-gray-900 focus:border-blue-500"

  return (
    <div className={`min-h-screen transition-colors duration-300 ${themeClass} font-sans selection:bg-blue-500 selection:text-white`}>

      {/* Top Navigation Bar */}
      <nav className={`fixed w-full z-10 backdrop-blur-md border-b px-6 py-4 flex justify-between items-center transition-colors duration-300 ${darkMode ? "border-slate-800/80 bg-slate-900/80" : "border-gray-200/80 bg-white/80"}`}>
        <div className="flex items-center gap-3">
          <div className="p-2 bg-gradient-to-tr from-blue-600 to-indigo-600 rounded-lg shadow-lg shadow-blue-500/20">
            <Activity className="text-white" size={24} />
          </div>
          <div>
            <h1 className="text-xl font-bold tracking-tight">SubSpace <span className="text-blue-500">Automator</span></h1>
            <p className={`text-xs ${darkMode ? "text-slate-400" : "text-gray-500"}`}>v1.0.0 • Go-Rod Engine</p>
          </div>
        </div>

        <div className="flex items-center gap-4">
          <button
            onClick={() => setDarkMode(!darkMode)}
            className={`p-2 rounded-full transition-all ${darkMode ? "bg-slate-800 text-yellow-400 hover:bg-slate-700" : "bg-gray-100 text-gray-600 hover:bg-gray-200"}`}
          >
            {darkMode ? <Sun size={20} /> : <Moon size={20} />}
          </button>
        </div>
      </nav>

      <main className="pt-28 pb-12 px-6 max-w-7xl mx-auto space-y-8">

        {/* Statistics Overview (The Box logic requested) */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <StatCard
            title="Profiles Scraped"
            value={profilesFound}
            icon={<Users className="text-blue-500" />}
            trend="+12% today"
            darkMode={darkMode}
            cardClass={cardClass}
          />
          <StatCard
            title="Requests Sent"
            value={requestsSent}
            icon={<User className="text-emerald-500" />}
            trend="Active"
            darkMode={darkMode}
            cardClass={cardClass}
          />
          <StatCard
            title="Messages"
            value={messagesSent}
            icon={<MessageSquare className="text-purple-500" />}
            trend="Pending"
            darkMode={darkMode}
            cardClass={cardClass}
          />
        </div>

        {/* Control Center */}
        <div className={`p-6 rounded-2xl border backdrop-blur-sm shadow-xl ${cardClass} transition-all`}>
          <div className="flex flex-col md:flex-row justify-between items-end gap-6">

            <div className="w-full space-y-4">
              <div className="flex justify-between items-center">
                <h2 className="text-lg font-semibold flex items-center gap-2">
                  <Search size={18} className="text-blue-500" /> Target Configuration
                </h2>
                <button
                  onClick={() => setShowCreds(!showCreds)}
                  className="text-xs text-blue-500 hover:underline flex items-center gap-1"
                >
                  <Lock size={12} /> {showCreds ? "Use Default .env" : "Set Custom Credentials"}
                </button>
              </div>

              {/* Dynamic Credentials inputs with smooth animation */}
              <div className={`grid gap-3 transition-all duration-300 overflow-hidden ${showCreds ? "max-h-20 opacity-100 mb-2" : "max-h-0 opacity-0 mb-0"}`}>
                <div className="flex gap-4">
                  <input
                    type="email"
                    placeholder="LinkedIn Email"
                    value={creds.email}
                    onChange={e => setCreds({ ...creds, email: e.target.value })}
                    className={`w-1/2 px-4 py-2 text-sm rounded-lg border outline-none transition-all ${inputClass}`}
                  />
                  <input
                    type="password"
                    placeholder="LinkedIn Password"
                    value={creds.password}
                    onChange={e => setCreds({ ...creds, password: e.target.value })}
                    className={`w-1/2 px-4 py-2 text-sm rounded-lg border outline-none transition-all ${inputClass}`}
                  />
                </div>
              </div>

              <div className="flex gap-3">
                <div className="relative flex-grow">
                  <input
                    type="text"
                    value={searchQuery}
                    onChange={(e) => setSearchQuery(e.target.value)}
                    disabled={status.running}
                    className={`w-full pl-10 pr-4 py-3 rounded-xl border outline-none transition-all font-medium ${inputClass} ${status.running ? "opacity-50 cursor-not-allowed" : ""}`}
                    placeholder="Ex: 'Software Engineer', 'Recruiter', 'Founder'..."
                  />
                  <Search className="absolute left-3 top-3.5 text-gray-400" size={18} />
                </div>

                {status.running ? (
                  <button
                    onClick={stopAutomation}
                    className="px-8 py-3 bg-red-500 hover:bg-red-600 text-white rounded-xl font-bold flex items-center gap-2 shadow-lg shadow-red-500/30 transition-all hover:scale-105 active:scale-95"
                  >
                    <Square size={18} fill="currentColor" /> Stop
                  </button>
                ) : (
                  <button
                    onClick={startAutomation}
                    className="px-8 py-3 bg-blue-600 hover:bg-blue-700 text-white rounded-xl font-bold flex items-center gap-2 shadow-lg shadow-blue-500/30 transition-all hover:scale-105 active:scale-95"
                  >
                    <Play size={18} fill="currentColor" /> Start Automation
                  </button>
                )}
              </div>
            </div>

          </div>
        </div>

        {/* Data & Logs Grid */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">

          {/* Live Logs Terminal */}
          <div className="lg:col-span-1 space-y-4">
            <h3 className="font-semibold flex items-center gap-2"><shield className="text-green-500" size={18} /> System Logs</h3>
            <div className={`h-[500px] rounded-2xl p-4 font-mono text-xs overflow-y-auto custom-scrollbar border ${darkMode ? "bg-black border-slate-800 text-green-400" : "bg-gray-900 border-gray-800 text-green-300"}`}>
              {status.logs.length === 0 ? (
                <div className="flex flex-col items-center justify-center h-full text-gray-600 gap-2">
                  <Database size={32} />
                  <p>Waiting for process...</p>
                </div>
              ) : (
                status.logs.slice().reverse().map((log, i) => (
                  <div key={i} className="mb-2 break-words leading-relaxed">
                    <span className="opacity-50 mr-2">{log.split(']')[0]}]</span>
                    <span>{log.split(']')[1] || log}</span>
                  </div>
                ))
              )}
            </div>
          </div>

          {/* Results Table */}
          <div className="lg:col-span-2 space-y-4">
            <div className="flex justify-between items-center">
              <h3 className="font-semibold flex items-center gap-2"><Database className="text-purple-500" size={18} /> Scraped Data</h3>
              <span className={`text-xs px-2 py-1 rounded-md ${darkMode ? "bg-slate-800 text-slate-400" : "bg-gray-200 text-gray-600"}`}>Live Feed</span>
            </div>

            <div className={`rounded-2xl border overflow-hidden ${cardClass}`}>
              <div className="overflow-x-auto">
                <table className="w-full text-left">
                  <thead>
                    <tr className={`text-xs uppercase tracking-wider ${darkMode ? "bg-slate-800/50 text-slate-400" : "bg-gray-50 text-gray-500"}`}>
                      <th className="px-6 py-4">Status</th>
                      <th className="px-6 py-4">Candidate Profile</th>
                      <th className="px-6 py-4">Timestamp</th>
                    </tr>
                  </thead>
                  <tbody className="divide-y divide-gray-100/10">
                    {activities.length === 0 ? (
                      <tr>
                        <td colSpan="3" className="px-6 py-12 text-center text-gray-500">
                          No data found yet. Start an automation to begin scraping.
                        </td>
                      </tr>
                    ) : (
                      activities.map((act) => (
                        <tr key={act.ID} className={`group transition-colors ${darkMode ? "hover:bg-slate-800/30" : "hover:bg-blue-50/50"}`}>
                          <td className="px-6 py-4 whitespace-nowrap">
                            <StatusBadge action={act.Action} />
                          </td>
                          <td className="px-6 py-4">
                            <div className="flex flex-col">
                              <span className="font-medium">{act.Metadata}</span>
                              <a href={act.ProfileURL} target="_blank" className="text-xs text-blue-500 hover:underline truncate max-w-xs">{act.ProfileURL}</a>
                            </div>
                          </td>
                          <td className="px-6 py-4 text-xs opacity-60">
                            {new Date(act.Timestamp).toLocaleTimeString()}
                          </td>
                        </tr>
                      ))
                    )}
                  </tbody>
                </table>
              </div>
            </div>
          </div>

        </div>

      </main>
    </div>
  )
}

// Helper Components

function StatCard({ title, value, icon, trend, darkMode, cardClass }) {
  return (
    <div className={`p-6 rounded-2xl border shadow-sm flex items-center justify-between ${cardClass}`}>
      <div>
        <p className={`text-sm font-medium ${darkMode ? "text-slate-400" : "text-gray-500"}`}>{title}</p>
        <h3 className="text-3xl font-bold mt-1">{value}</h3>
      </div>
      <div className={`p-3 rounded-xl ${darkMode ? "bg-slate-800" : "bg-gray-100"}`}>
        {icon}
      </div>
    </div>
  )
}

function StatusBadge({ action }) {
  if (action === 'SEARCH_FOUND') return <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">Scraped</span>
  if (action === 'CONNECT') return <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">Req Sent</span>
  if (action === 'MESSAGE') return <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-purple-100 text-purple-800">Messaged</span>
  return <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-gray-100 text-gray-800">{action}</span>
}

export default App
