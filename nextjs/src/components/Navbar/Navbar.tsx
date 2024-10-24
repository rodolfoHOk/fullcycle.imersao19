import Link from 'next/link';
import { SearchBar } from './SearchBar';

export function Navbar() {
  return (
    <header className="bg-primary border-secondary border-b-2">
      <div className="container mx-auto px-4 py-2 flex items-center justify-between">
        <div className="flex items-center space-x-4">
          <Link href="/" className="text-2xl font-bold text-[#ffcd00]">
            FCTube
          </Link>
        </div>

        <div className="w-1/2 relative">
          <SearchBar />
        </div>

        <div className="flex items-center space-x-4">
          <a
            href="#"
            className="text-primary hover:text-primary/85 transition-colors"
          >
            Login
          </a>
        </div>
      </div>
    </header>
  );
}
