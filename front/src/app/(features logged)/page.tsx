"use client";
import { GetProfile } from "@/lib/services/profile-service";

export default function Home() {
  void GetProfile().then();

  return (
    <div className="grid grid-cols-1 grid-rows-6 gap-4">
      {/* header */}
      <section className="grid grid-cols-4 grid-rows-3 gap-4">
        <div className="col-span-2">1</div>
        <div className="col-span-2 col-start-3">2</div>
        <div className="col-span-4 row-start-2">3</div>
        <div className="row-start-3">4</div>
        <div className="row-start-3">5</div>
        <div className="row-start-3">6</div>
        <div className="row-start-3">7</div>
      </section>

      {/* Presente */}
      <section>
        <div className="grid grid-cols-2 grid-rows-3 gap-4">
          <div className="col-span-2">11</div>
          <div className="row-start-2">12</div>
          <div className="row-start-2">13</div>
          <div>14</div>
          <div className="row-start-3">15</div>
        </div>
      </section>
      {/* Passado */}
      <section>
        <div className="grid grid-cols-4 grid-rows-7 gap-4">
          <div className="col-span-4">16</div>
          <div className="col-span-4 row-start-2">17</div>
          <div className="col-span-2 row-start-3">18</div>
          <div className="col-span-2 col-start-3 row-start-3">19</div>
          <div className="col-span-2 row-start-4">20</div>
          <div className="col-span-2 col-start-3 row-start-4">21</div>
          <div className="col-span-3 row-start-5">22</div>
          <div className="col-start-4 row-start-5">23</div>
          <div className="col-span-3 row-start-6">24</div>
          <div className="col-start-4 row-start-6">25</div>
          <div className="col-span-4 row-start-7">26</div>
        </div>
      </section>
      {/* Construção */}
      <section>
        <div className="grid grid-cols-2 grid-rows-2 gap-4">
          <div className="col-span-2">27</div>
          <div className="row-start-2">28</div>
          <div className="row-start-2">29</div>
        </div>
      </section>
      {/* Futuro */}
      <section>
        <div className="grid grid-cols-2 grid-rows-2 gap-4">
          <div className="col-span-2">27</div>
          <div className="row-start-2">28</div>
          <div className="row-start-2">29</div>
        </div>
      </section>
      <section>6</section>
    </div>
  );
}
