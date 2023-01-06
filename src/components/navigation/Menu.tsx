import type { Component } from 'solid-js'
import { cx } from '../util/classes'

export interface MenuProps {
  visible: boolean
  baseSite: string
  blogSite: string
  github: string
}

export const menuId = 'menu'
export const closeMenuId = 'close-menu'
export const menuVisibleCx = ['right-0']
export const menuInvisibleCx = ['-right-96']

export const Menu: Component<MenuProps> = ({
  visible,
  baseSite,
  blogSite,
  github,
}) => {
  return (
    <div
      id={menuId}
      role="navigation"
      class={cx([
        visible ? menuVisibleCx : menuInvisibleCx,
        'text-white',
        'flex',
        'bg-neutral-800',
        'fixed',
        'top-0',
        'h-full',
        'w-96',
        'p-12',
        'box-border',
        'z-50',
        'transition-right',
      ])}
    >
      <span
        id={closeMenuId}
        class={cx(['flex', 'cursor-pointer', 'absolute', 'top-12', 'left-12'])}
      >
        <img src="/close.svg" alt="close" />
      </span>
      <div class="my-auto">
        <h2 class="text-lg font-geometric">Where do you want to go?</h2>
        <menu role="menu" id="menu-links">
          <li>
            <a href={baseSite} title="Home">
              home
            </a>
          </li>
          <li>
            <a href={blogSite} title="Blog">
              blog
            </a>
          </li>
          <li>
            <a href={`${baseSite}#me`} title="About me">
              about me
            </a>
          </li>
          <li>
            <a href={github} title="My Github">
              code
            </a>
          </li>
        </menu>
      </div>
    </div>
  )
}
