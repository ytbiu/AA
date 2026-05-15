'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'

export default function Header() {
  const pathname = usePathname()

  const navItems = [
    { name: '首页', path: '/' },
    { name: '水龙头', path: '/faucet' },
    { name: '7702授权', path: '/authorize' },
    { name: '清除授权', path: '/clear' },
    { name: 'USDT转账', path: '/transfer' },
    { name: '管理', path: '/admin' },
  ]

  return (
    <header className="bg-white shadow-sm border-b">
      <nav className="container mx-auto px-4">
        <div className="flex items-center justify-between h-16">
          <Link href="/" className="text-xl font-bold text-blue-600">
            AA Wallet
          </Link>
          <div className="flex gap-4">
            {navItems.map((item) => (
              <Link
                key={item.path}
                href={item.path}
                className={`px-3 py-2 rounded-md text-sm font-medium transition-colors ${
                  pathname === item.path
                    ? 'bg-blue-100 text-blue-700'
                    : 'text-gray-600 hover:bg-gray-100'
                }`}
              >
                {item.name}
              </Link>
            ))}
          </div>
        </div>
      </nav>
    </header>
  )
}